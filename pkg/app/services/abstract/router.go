package abstract

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	RouteRoot() *gin.Engine
}
