package session

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
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

func (m *Middleware) loadConfig() {
	m.loadDefaults()
	m.reload()
}

func (m *Middleware) ConfigGroup() abstract.ConfigGroup {
	return m.cfg.GetConfigGroup(m.Name())
}

func (m *Middleware) loadDefaults() {
	for key, def := range map[string]string{
		keyBackend:        defaultBackend,
		keyCorsSameSite:   fmt.Sprintf("%v", defaultCorsSameSite),
		keyMemstoreSecret: defaultMemstoreSecret,
		keyRedisHostname:  defaultRedisHostname,
		keyRedisPort:      fmt.Sprintf("%v", defaultRedisPort),
		keyRedisPassword:  defaultRedisPassword,
		keyRedisSize:      fmt.Sprintf("%v", defaultRedisSize),
		keyRedisProtocol:  defaultRedisProtocol,
		keyRedisSecret:    defaultRedisSecret,
	} {
		m.ConfigGroup().Default(key, def)
	}
}

func (m *Middleware) handleConfigChanges() {
	for {
		m.updateConfig()
		time.Sleep(time.Second)
	}
}

func (m *Middleware) reload() {
	m.config.backend = m.cfg.Get(m.Name(), keyBackend)
	m.config.memstore.secret = m.cfg.Get(m.Name(), keyMemstoreSecret)
	m.config.redis.hostname = m.cfg.Get(m.Name(), keyRedisHostname)
	m.config.redis.protocol = m.cfg.Get(m.Name(), keyRedisProtocol)
	m.config.redis.password = m.cfg.Get(m.Name(), keyRedisPassword)
	m.config.redis.secret = m.cfg.Get(m.Name(), keyRedisSecret)
	m.config.redis.port = s2i(m.cfg.Get(m.Name(), keyRedisPort), defaultRedisPort)
	m.config.redis.size = s2i(m.cfg.Get(m.Name(), keyRedisSize), defaultRedisSize)
	m.config.cors.sameSite = http.SameSite(s2i(m.cfg.Get(m.Name(), keyCorsSameSite), int(defaultCorsSameSite)))
}

func (m *Middleware) updateConfig() {
	if m.sessionMiddleware != nil {
		return
	}

	if err := m.initSessionsMiddleware(); err != nil {
		m.log.Fatal().Msgf("[%s] initializing session middleware: %v", m.Name(), err)
	}
}

func s2i(s string, def int) int {
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return int(i)
	}

	return def
}
