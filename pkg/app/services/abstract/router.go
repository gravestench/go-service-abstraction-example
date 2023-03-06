package abstract

import (
	"github.com/gin-gonic/gin"
)

// Router is just responsible for yielding the root route handler.
// Services will use this in order to set up their own routes.
type Router interface {
	RouteRoot() *gin.Engine
}
