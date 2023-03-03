package webserver

import (
	"fmt"
	"net/http"
)

func (s *Service) initHttpServer() {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.config.port),
		Handler: s.router.RouteRoot(),
	}

	if err := s.server.ListenAndServe(); err != nil {
	}
}
