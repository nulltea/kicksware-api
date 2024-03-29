package server

import (
	"context"
	"crypto/rsa"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang/glog"
	grpcLog "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"

	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"

	"go.kicksware.com/api/shared/api/gRPC"
	"go.kicksware.com/api/shared/api/jwt"
	"go.kicksware.com/api/shared/core"
	"go.kicksware.com/api/shared/core/meta"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"go.elastic.co/apm/module/apmgrpc"
	"go.elastic.co/apm/module/apmhttp"

	"github.com/soheilhy/cmux"
)

type instance struct {
	Address  string
	Gateway  cmux.CMux
	REST     *http.Server
	GRPC     *grpc.Server
	AMQP     core.Handler
	TLS      credentials.TransportCredentials
	Auth     *gRPC.AuthServerInterceptor
	Logger   *logrus.Logger
	LogEntry *logrus.Entry
}

func NewInstance(addr string) core.Server {
	return &instance{
		Address: addr,
	}
}

func (s *instance) SetupEncryption(cert *meta.TLSCertificate) {
	if cert.EnableTLS {
		cred, err := gRPC.LoadServerTLSCredentials(cert); if err != nil {
			glog.Fatalln("cannot load TLS credentials: ", err)
		}
		s.TLS = cred
	}
}

func (s *instance) SetupAuth(pb *rsa.PublicKey, accessRoles meta.AccessConfig) {
	JWTManager := &jwt.TokenManager{
		PublicKey: pb,
	}
	s.Auth = gRPC.NewAuthServerInterceptor(JWTManager, accessRoles)
}

func (s *instance) SetupLogger() {
	s.Logger = &logrus.Logger{
		Level: logrus.InfoLevel,
	}
	s.LogEntry = logrus.NewEntry(s.Logger)
	grpcLog.ReplaceGrpcLogger(s.LogEntry)
}

func (s *instance) SetupREST(router chi.Router) {
	s.REST = &http.Server{
		Addr: s.Address,
	}
	s.REST.Handler = apmhttp.Wrap(router);
}

func (s *instance) SetupGRPC(fn func(srv *grpc.Server)) {
	unaryInterceptors := []grpc.UnaryServerInterceptor {
		apmgrpc.NewUnaryServerInterceptor(),
		grpcTags.UnaryServerInterceptor(),
		grpcRecovery.UnaryServerInterceptor(),
	}
	streamInterceptors := []grpc.StreamServerInterceptor {
		grpcTags.StreamServerInterceptor(),
		grpcRecovery.StreamServerInterceptor(),
	}
	options := []grpc.ServerOption{
		grpc.MaxSendMsgSize(25 * 1024 * 1024),
		grpc.MaxRecvMsgSize(25 * 1024 * 1024),
	}; if s.TLS != nil {
		options = append(options, grpc.Creds(s.TLS))
	}; if s.Auth != nil {
		unaryInterceptors = append(unaryInterceptors, s.Auth.Unary())
		streamInterceptors = append(streamInterceptors, s.Auth.Stream())
	}; if s.LogEntry != nil {
		unaryInterceptors = append(unaryInterceptors, grpcLog.UnaryServerInterceptor(s.LogEntry))
		streamInterceptors = append(streamInterceptors, grpcLog.StreamServerInterceptor(s.LogEntry))
	}

	options = append(options,
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(unaryInterceptors...)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(streamInterceptors...)),
	)

	s.GRPC = grpc.NewServer(options...)

	fn(s.GRPC)
	reflection.Register(s.GRPC)
}

func (s *instance) SetupAMQP(handler core.Handler) {
	s.AMQP = handler
}

func (s *instance) Start() {
	errs := make(chan error, 2)

	lstn, err := net.Listen("tcp", s.Address); if err == nil {
		fmt.Println(fmt.Sprintf("Microservice launched to address http://%v", s.Address))
	} else {
		glog.Fatalf("Failed to listen on %v: %q", s.Address, err)
	}

	s.Gateway = cmux.New(lstn)
	grpcL := s.Gateway.Match(cmux.HTTP2()); if s.TLS != nil {
		grpcL = s.Gateway.Match(cmux.TLS())
	}
	restL := s.Gateway.Match(cmux.HTTP1Fast())

	if s.REST != nil {
		go func() {
			errs <- s.REST.Serve(restL)
		}()
	}

	if s.GRPC != nil {
		go func() {
			errs <- s.GRPC.Serve(grpcL)
		}()
	}

	if s.AMQP != nil {
		s.AMQP.Setup()
		go func() {
			errs <- s.AMQP.Serve()
		}()
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errs <- fmt.Errorf("%s", <-c)
		s.Shutdown()
	}()

	s.Gateway.Serve()

	fmt.Printf("Terminated %s", <-errs)
}

func (s *instance) Shutdown() {
	if s.REST != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.REST.Shutdown(ctx)
		if err != nil {
			errors.Wrap(err, "Failed to shutdown rest server gracefully")
		}
	}

	if s.GRPC != nil {
		s.GRPC.GracefulStop()
	}
}
