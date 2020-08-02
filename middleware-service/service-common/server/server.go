package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang/glog"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
	"github.com/soheilhy/cmux"
)

type instance struct {
	Address string
	REST *http.Server
	GRPC *grpc.Server
	Gateway cmux.CMux
}

func NewInstance(addr string) core.Server {
	return &instance{
		REST: &http.Server{
			Addr:      addr,
		},
		GRPC: grpc.NewServer(
			grpc.MaxSendMsgSize(20 * 1024 * 1024),
		),
		Address: addr,
	}
}

func (s *instance) SetupREST(router chi.Router) {
	s.REST.Handler = router;
}

func (s *instance) SetupRoutes(router chi.Router) {
	s.SetupREST(router)
}

func (s *instance) SetupGRPC(fn func(srv *grpc.Server)) {
	fn(s.GRPC)
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

	go func() {
		errs <- s.REST.Serve(restL)
	}()

	go func() {
		errs <- s.GRPC.Serve(grpcL)
	}()

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
