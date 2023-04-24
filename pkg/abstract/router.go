package abstract

import (
	"github.com/gin-gonic/gin"
)

// Router is just responsible for yielding the root route handler.
// Services will use this in order to set up their own routes.
type Router interface {
	RouteRoot() *gin.Engine
	RouteProtected() *gin.RouterGroup
}

// HasRouteSlug describes a service that has an identifier that is used
// as a prefix for its subroutes
type HasRouteSlug interface {
	Slug() string
}

// RouteInitializer is a type of service that will
// set up its own web routes using a base route group
type RouteInitializer interface {
	Service
	InitRoutes(*gin.RouterGroup) error
}

// ProtectedRouteInitializer is just like a RouteInitializer,
// but it uses a route group with session management
type ProtectedRouteInitializer interface {
	Service
	InitProtectedRoutes(*gin.RouterGroup) error
}
