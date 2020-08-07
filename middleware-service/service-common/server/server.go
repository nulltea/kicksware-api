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
	"github.com/pkg/errors"
	"github.com/timoth-y/kicksware-platform/middleware-service/user-service/core/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/core"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/service/gRPC"
	"github.com/timoth-y/kicksware-platform/middleware-service/service-common/service/jwt"

	"github.com/soheilhy/cmux"
)

type instance struct {
	Address string
	Gateway cmux.CMux
	REST *http.Server
	GRPC *grpc.Server
	Auth *gRPC.AuthServerInterceptor
}

func NewInstance(addr string) core.Server {
	return &instance{
		Address: addr,
	}
}

func (s *instance) SetupAuth(pb *rsa.PublicKey, accessRoles map[string][]model.UserRole) {
	jwtManager := &jwt.TokenManager{
		PublicKey: pb,
	}
	s.Auth = gRPC.NewAuthServerInterceptor(jwtManager, accessRoles)
}

func (s *instance) SetupREST(router chi.Router) {
	s.REST = &http.Server{
		Addr: s.Address,
	}
	s.REST.Handler = router;
}

func (s *instance) SetupGRPC(fn func(srv *grpc.Server)) {
	options := []grpc.ServerOption{
		grpc.MaxSendMsgSize(25 * 1024 * 1024),
		grpc.MaxRecvMsgSize(25 * 1024 * 1024),
	}; if s.Auth != nil {
		options = append(options, grpc.UnaryInterceptor(s.Auth.Unary()), grpc.StreamInterceptor(s.Auth.Stream()))
	}

	s.GRPC = grpc.NewServer(options...)

	fn(s.GRPC)
	reflection.Register(s.GRPC)
}

func (s *instance) Start() {
	errs := make(chan error, 2)

	lstn, err := net.Listen("tcp", s.Address); if err == nil {
		fmt.Println(fmt.Sprintf("Microservice launched to address http://%v", s.Address))
	} else {
		glog.Fatalf("Failed to listen on %v: %q", s.Address, err)
	}

	s.Gateway = cmux.New(lstn)
	grpcL := s.Gateway.Match(cmux.HTTP2())
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