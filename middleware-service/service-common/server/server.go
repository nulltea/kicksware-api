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
	"google.golang.org/grpc"
	"github.com/pkg/errors"

	"github.com/timoth-y/sneaker-resale-platform/middleware-service/service-common/core"
)

type instance struct {
	Listener *net.Listener
	Address string
	REST *http.Server
	GRPC *grpc.Server
}

func NewInstance(addr string) core.Server {
	return &instance{
		REST: &http.Server{
			Addr:      addr,
		},
		GRPC: grpc.NewServer(),
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

	if lstn, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080)); err == nil {
		s.Listener = &lstn
		fmt.Println(fmt.Sprintf("Microservice launched to address http://%v", s.Address))
	} else {
		glog.Fatalf("Failed to listen on %v: %q", s.Address, err)
	}

	go func() {
		errs <- s.REST.Serve(*s.Listener)
	}()

	go func() {
		errs <- s.GRPC.Serve(*s.Listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errs <- fmt.Errorf("%s", <-c)
		s.Shutdown()
	}()

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