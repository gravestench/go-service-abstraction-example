package cors

import (
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
)

const (
	keyOrigins     = "origins"
	defaultOrigins = "https://localhost:4200"

	keyMethods     = "methods"
	defaultMethods = "GET,POST,PUT,PATCH,DELETE,OPTIONS"

	keyHeaders     = "headers"
	defaultHeaders = "Content-Type"

	keyAllowCredentials     = "allowcredentials"
	defaultAllowCredentials = "true"

	keySameSite     = "samesite"
	defaultSameSite = "4" // http.SameSiteNoneMode
)

func New(log abstract.Logger, cfg abstract.ConfigurationManager) *Middleware {
	m := &Middleware{
		log: log,
		cfg: cfg,
	}

	m.init()

	return m
}

type Middleware struct {
	log    abstract.Logger
	cfg    abstract.ConfigurationManager
	router abstract.Router
	config *config
}

type config struct {
	origins          []string
	methods          []string
	headers          []string
	allowCredentials string
	sameSite         string
}

func (m *Middleware) init() {
	m.loadDefaults()
	m.loadConfig()
}

func (m *Middleware) Name() string {
	return "CORS Middleware"
}

func (m *Middleware) getConfig() cors.Config {
	return cors.Config{
		AllowOrigins:     m.config.origins,
		AllowMethods:     m.config.methods,
		AllowHeaders:     m.config.headers,
		AllowCredentials: m.config.allowCredentials == "true",
	}
}

func (m *Middleware) GetCorsHandler() gin.HandlerFunc {
	if m.config == nil {
		return nil
	}

	return cors.New(m.getConfig())
}

func (m *Middleware) loadDefaults() {
	group := m.cfg.GetConfigGroup(m.Name())

	group.Default(keyOrigins, defaultOrigins)
	group.Default(keyMethods, defaultMethods)
	group.Default(keyHeaders, defaultHeaders)
	group.Default(keyAllowCredentials, defaultAllowCredentials)
	group.Default(keySameSite, defaultSameSite)
}

func (m *Middleware) loadConfig() {
	m.config = &config{}
	group := m.cfg.GetConfigGroup(m.Name())

	m.config.origins = strings.Split(group.Get(keyOrigins), ",")
	m.config.methods = strings.Split(group.Get(keyMethods), ",")
	m.config.headers = strings.Split(group.Get(keyHeaders), ",")
	m.config.allowCredentials = group.Get(keyAllowCredentials)
	m.config.sameSite = group.Get(keySameSite)
}
