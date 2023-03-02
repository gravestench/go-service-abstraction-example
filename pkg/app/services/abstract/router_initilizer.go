package abstract

import (
	"github.com/gin-gonic/gin"
)

type RouteInitializer interface {
	Service
	InitRoutes(*gin.RouterGroup) error
	Slug() string
}
