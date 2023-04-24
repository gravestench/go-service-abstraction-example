package router

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) loadConfig() {
	strDebug := s.cfgManager.Get(s.Name(), "debug")

	if strDebug == "true" {
		s.config.debug = true
		return
	}

	s.cfgManager.Set(s.Name(), "debug", "false")
}

func (s *Service) handleConfigChanges() {
	for {
		s.handleIsDebug()
		time.Sleep(time.Millisecond * 10)
	}
}

func (s *Service) handleIsDebug() {
	strIsDebug := s.cfgManager.Get(s.Name(), "debug")
	isDebug := strIsDebug == "true"

	if s.config.debug != isDebug {
		s.log.Info().Msgf("setting debug mode to %v", strIsDebug)
		s.config.debug = isDebug

		if s.config.debug {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}
	}
}
