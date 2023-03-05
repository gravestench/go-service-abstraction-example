package session_middleware

import (
	"fmt"

	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
)

type sessionStorageType = string

const (
	sessionMemstore sessionStorageType = "memstore"
	sessionRedis    sessionStorageType = "redis"
)

func (s *Service) initSessionsMiddleware() error {
	s.setSameSite(s.config.cors.sameSite)

	current := fmt.Sprintf("%+v", s.config)

	if s.previousConfig == current {
		return nil
	}

	if s.previousConfig == "" {
		s.log.Info().Msgf("[%s] initializing", s.Name())
	} else {
		s.log.Info().Msgf("[%s] re-initializing", s.Name())
	}

	s.previousConfig = current

	switch s.config.backend {
	case sessionMemstore:
		s.log.Info().Msgf("[%s] using memstore for session management", s.Name())

		store := memstore.NewStore([]byte(s.config.memstore.secret))
		s.sessionMiddleware = s.newSessionHandler("session", store)
	case sessionRedis:
		s.log.Info().Msgf("[%s] using redis for session management", s.Name())

		size, protocol := s.config.redis.size, s.config.redis.protocol
		address := fmt.Sprintf("%s:%d", s.config.redis.hostname, s.config.redis.port)
		pw, secret := s.config.redis.password, []byte(s.config.redis.secret)

		store, redisErr := redis.NewStore(size, protocol, address, pw, secret)
		if redisErr != nil {
			return fmt.Errorf("could not create redis store: %v", redisErr)
		}

		s.sessionMiddleware = s.newSessionHandler("session", store)
	default:
		return fmt.Errorf("unknown session backend specified: %v", s.config.backend)
	}

	s.router.Use(s.sessionMiddleware)
	s.router.Use(Checkauth())
	s.router.Use(CheckRole())

	return nil
}
