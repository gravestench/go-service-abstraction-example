package session

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/abstract"
	"github.com/gravestench/go-service-abstraction-example/pkg/abstract/models"
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
	log abstract.Logger
	cfg abstract.ConfigurationManager

	sessionMiddleware gin.HandlerFunc
	config
	previousConfig string // to check when we need to reload
}

func (m *Middleware) init() {
	m.config.sessionOptions.HttpOnly = true
	m.config.sessionOptions.SameSite = http.SameSiteNoneMode
	m.config.sessionOptions.Secure = true
	m.config.sessionOptions.Path = "/"
	m.config.sessionOptions.MaxAge = secondsInDay * daysInWeek

	m.loadConfig()
	go m.handleConfigChanges()
}

func (m *Middleware) Name() string {
	return "Session Middleware"
}

const (
	daysInWeek   int = 7
	secondsInDay int = 86400
)

type UserSession struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (m *Middleware) setSameSite(level http.SameSite) {
	switch level {
	case http.SameSiteLaxMode, http.SameSiteStrictMode, http.SameSiteNoneMode:
		m.sessionOptions.SameSite = level
	default:
		m.sessionOptions.SameSite = http.SameSiteStrictMode
	}

	m.cfg.Set(m.Name(), keyCorsSameSite, m.config.cors.sameSite)
}

func (m *Middleware) SessionHandler() gin.HandlerFunc {
	return m.sessionMiddleware
}

func (m *Middleware) initSessionHandler(name string, store sessions.Store) gin.HandlerFunc {
	gob.Register(&models.User{})
	store.Options(m.sessionOptions)

	return sessions.Sessions(name, store)
}

func (m *Middleware) Checkauth(c *gin.Context) {
	var unguardedRoutes = map[string]struct{}{
		"POST/api/session/login":  {},
		"GET/api/session/logout":  {},
		"GET/api/session/exists":  {},
		"GET/api/ui-config/terms": {},
	}

	var str strings.Builder
	str.WriteString(c.Request.Method)
	str.WriteString(c.Request.URL.Path)

	if _, ok := unguardedRoutes[str.String()]; ok {
		return
	}
	session := sessions.Default(c)

	val := session.Get("user")
	_, ok := val.(*models.User)
	if !ok {
		c.AbortWithStatus(http.StatusForbidden)
	}
}

func (m *Middleware) CheckRole(c *gin.Context) {
	// get the user session
	session := sessions.Default(c)
	rawUser := session.Get("user")

	if rawUser == nil {
		// nothing to do if user is not logged in
		return
	}

	// this shouldn't explode, but we need to catch possible nil pointer exception here
	userModel, ok := rawUser.(*models.User)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	switch userModel.Role {
	case models.UserRoleLimitedUser:
		handleLimitedUserRole(userModel, c)
	}
}

func handleLimitedUserRole(u *models.User, c *gin.Context) {
	switch c.Request.Method {
	case "PUT":
		if isLimitedUserSettingOwnPassword(u, c) {
			return
		}
		c.AbortWithStatus(http.StatusForbidden)
	case "POST":
		c.AbortWithStatus(http.StatusForbidden)
	case "DELETE":
		c.AbortWithStatus(http.StatusForbidden)
	case "GET":
		return
	}
}

func isLimitedUserSettingOwnPassword(u *models.User, c *gin.Context) bool {
	// this is hard-coded to match the api route for user management,
	// see pkg/app/route_constants.go
	isUserManagementRoute := strings.HasSuffix(c.Request.URL.String(), "users/profile")
	if !isUserManagementRoute {
		return false
	}

	byteBody, byteBodyErr := ioutil.ReadAll(c.Request.Body)
	if byteBodyErr != nil {
		return false
	}

	var userObject models.User
	jsonErr := json.Unmarshal(byteBody, &userObject)
	if jsonErr != nil {
		return false
	}

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))

	if userObject.Name != u.Name {
		return false
	}

	return true
}
