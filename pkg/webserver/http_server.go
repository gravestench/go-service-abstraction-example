package webserver

import (
	"fmt"
	"net/http"
)

func (s *Service) initHttpServer() {
	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%v", s.Config().Get(keyPort)),
		Handler: s.router.RouteRoot(),
	}

	if err := s.server.ListenAndServe(); err != nil {
		s.log.Fatal().Msgf("initializing server: %v", err)
	}
}
