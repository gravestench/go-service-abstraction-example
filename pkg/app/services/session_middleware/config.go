package session_middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
)

const (
	keyBackend     = "backend"
	defaultBackend = "memstore"
)

const (
	keyCorsSameSite     = "cors.sameSite"
	defaultCorsSameSite = http.SameSiteStrictMode
)

const ( // memstore
	keyMemstoreSecret     = "memstore.secret"
	defaultMemstoreSecret = "secret"
)

const ( // redis
	keyRedisHostname     = "redis.hostname"
	defaultRedisHostname = "localhost"

	keyRedisPort     = "redis.port"
	defaultRedisPort = 6379

	keyRedisPassword     = "redis.password"
	defaultRedisPassword = ""

	keyRedisSize     = "redis.size"
	defaultRedisSize = 10

	keyRedisProtocol     = "redis.protocol"
	defaultRedisProtocol = "tcp"

	keyRedisSecret     = "redis.secret"
	defaultRedisSecret = "secret"
)

type config struct {
	backend        string
	redis          redisConfig
	memstore       memstoreConfig
	cors           corsConfig // cross origin resource sharing
	sessionOptions sessions.Options
}

type redisConfig struct {
	hostname string
	port     int
	protocol string
	password string
	size     int
	secret   string
}

type memstoreConfig struct {
	secret string
}

type corsConfig struct {
	sameSite http.SameSite
}

func (s *Service) loadConfig() {
	keys := s.cfg.GetKeys(s.Name())

	if keys == nil || len(keys) == 0 {
		s.loadDefaults()
		return
	}

	s.reload()
}

func (s *Service) loadDefaults() {
	s.config.backend = defaultBackend
	s.cfg.Set(s.Name(), keyBackend, s.config.backend)

	s.config.memstore.secret = defaultMemstoreSecret
	s.cfg.Set(s.Name(), keyMemstoreSecret, s.config.memstore.secret)

	s.config.redis.hostname = defaultRedisHostname
	s.config.redis.port = defaultRedisPort
	s.config.redis.protocol = defaultRedisProtocol
	s.config.redis.password = defaultRedisPassword
	s.config.redis.size = defaultRedisSize
	s.config.redis.secret = defaultRedisSecret
	s.cfg.Set(s.Name(), keyRedisHostname, s.config.redis.hostname)
	s.cfg.Set(s.Name(), keyRedisProtocol, s.config.redis.port)
	s.cfg.Set(s.Name(), keyRedisPassword, s.config.redis.protocol)
	s.cfg.Set(s.Name(), keyRedisSecret, s.config.redis.password)
	s.cfg.Set(s.Name(), keyRedisPort, s.config.redis.size)
	s.cfg.Set(s.Name(), keyRedisSize, s.config.redis.secret)

	s.config.cors.sameSite = defaultCorsSameSite
	s.cfg.Set(s.Name(), keyCorsSameSite, s.config.cors.sameSite)
}

func (s *Service) handleConfigChanges() {
	for {
		s.updateConfig()
		time.Sleep(time.Second)
	}
}

func (s *Service) reload() {
	s.config.backend = s.cfg.Get(s.Name(), keyBackend)

	s.config.memstore.secret = s.cfg.Get(s.Name(), keyMemstoreSecret)

	s.config.redis.hostname = s.cfg.Get(s.Name(), keyRedisHostname)
	s.config.redis.protocol = s.cfg.Get(s.Name(), keyRedisProtocol)
	s.config.redis.password = s.cfg.Get(s.Name(), keyRedisPassword)
	s.config.redis.secret = s.cfg.Get(s.Name(), keyRedisSecret)
	s.config.redis.port = s2i(s.cfg.Get(s.Name(), keyRedisPort), defaultRedisPort)
	s.config.redis.size = s2i(s.cfg.Get(s.Name(), keyRedisSize), defaultRedisSize)

	s.config.cors.sameSite = http.SameSite(s2i(s.cfg.Get(s.Name(), keyCorsSameSite), int(defaultCorsSameSite)))
}

func (s *Service) updateConfig() {
	if s.sessionRouteRoot == nil {
		if err := s.initSessionsMiddleware(); err != nil {
			s.log.Fatal().Msgf("[%s] initializing session middleware: %v", s.Name(), err)
		}
	}
}

func s2i(s string, def int) int {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(i)
	}

	return def
}
