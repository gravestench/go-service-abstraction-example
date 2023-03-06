package abstract

import (
	"github.com/gin-gonic/gin"
)

// RouteInitializer is a type of service that will
// set up its own web routes using a base route group
type RouteInitializer interface {
	Service
	InitRoutes(*gin.RouterGroup) error
	Slug() string
}
