package session

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

func (m *Middleware) initSessionsMiddleware() error {
	m.setSameSite(m.config.cors.sameSite)

	current := fmt.Sprintf("%+v", m.config)

	if m.previousConfig == current {
		return nil
	}

	if m.previousConfig == "" {
		m.log.Info().Msgf("[%s] initializing", m.Name())
	} else {
		m.log.Info().Msgf("[%s] re-initializing", m.Name())
	}

	m.previousConfig = current

	switch m.config.backend {
	case sessionMemstore:
		m.log.Info().Msgf("[%s] using memstore for session management", m.Name())

		store := memstore.NewStore([]byte(m.config.memstore.secret))
		m.sessionMiddleware = m.initSessionHandler("session", store)
	case sessionRedis:
		m.log.Info().Msgf("[%s] using redis for session management", m.Name())

		size, protocol := m.config.redis.size, m.config.redis.protocol
		address := fmt.Sprintf("%s:%d", m.config.redis.hostname, m.config.redis.port)
		pw, secret := m.config.redis.password, []byte(m.config.redis.secret)

		store, redisErr := redis.NewStore(size, protocol, address, pw, secret)
		if redisErr != nil {
			return fmt.Errorf("could not create redis store: %v", redisErr)
		}

		m.sessionMiddleware = m.initSessionHandler("session", store)
	default:
		return fmt.Errorf("unknown session backend specified: %v", m.config.backend)
	}

	return nil
}
