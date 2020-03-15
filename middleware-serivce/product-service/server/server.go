package server

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"golang.org/x/crypto/acme/autocert"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

type instance struct {
	Server *http.Server
	Address string
	Host string
	CertManager *autocert.Manager
}

func NewInstance(addr, host string) *instance {
	cert := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(cacheDir()),
		HostPolicy: autocert.HostWhitelist(host),
	}
	return &instance{
		Server: &http.Server{
			Addr:      addr,
			TLSConfig: cert.TLSConfig(),
		},
		Address:     addr,
		Host:        host,
		CertManager: cert,
	}
}

func (s *instance) SetupRouter(router chi.Router) {
	s.Server.Handler = router;
}

func (s *instance) Start() {
	errs := make(chan error, 2)
	fmt.Println(fmt.Sprintf("Listening on port http/https://%v:%v", s.Address, s.Host))
	go func() {
		h := s.CertManager.HTTPHandler(nil)
		errs <- http.ListenAndServe(":http", h)
	}()

	go func() {
		errs <- s.Server.ListenAndServeTLS("", "")
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errs <- fmt.Errorf("%s", <-c)
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

func cacheDir() (dir string) {
	dir = filepath.Join(os.TempDir(), "cache-golang-autocert")
	if err := os.MkdirAll(dir, 0700); err == nil {
		return dir
	}
	return ""
}