package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"

	"service-common/core"
)

type instance struct {
	Server *http.Server
	Address string
}

func NewInstance(addr string) service.Server {
	return &instance{
		Server: &http.Server{
			Addr:      addr,
		},
		Address: addr,
	}
}

func (s *instance) SetupRoutes(router chi.Router) {
	s.Server.Handler = router;
}

func (s *instance) Start() {
	errs := make(chan error, 2)
	fmt.Println(fmt.Sprintf("Listening on port http://%v", s.Address))

	go func() {
		errs <- s.Server.ListenAndServe()
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
	if s.Server != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := s.Server.Shutdown(ctx)
		if err != nil {
			errors.Wrap(err, "Failed to shutdown rest server gracefully")
		}
	}
}