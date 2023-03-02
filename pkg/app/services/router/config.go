package router

import (
	"time"
)

func (s *Service) loadConfig() {
	time.Sleep(time.Millisecond * 10)

	strDebug := s.cfgManager.Get(s.Name(), "debug")

	if strDebug == "true" {
		s.config.debug = true
		return
	}

	s.cfgManager.Set(s.Name(), "debug", "false")
}
