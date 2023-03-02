package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "cfg"
}

func (s *Service) InitRoutes(group *gin.RouterGroup) error {
	group.GET(":group", s.getGroup)
	group.GET(":group/:key", s.getGroupKey)

	return nil
}

func (s *Service) getGroup(c *gin.Context) {
	groupKey := c.Param("group")

	if groupKey == "" {
		c.JSON(http.StatusBadRequest, "{}")
		return
	}

	c.JSON(http.StatusOK, s.store[groupKey])
}

func (s *Service) getGroupKey(c *gin.Context) {
	g := c.Param("group")
	k := c.Param("key")

	if g == "" {
		c.JSON(http.StatusBadRequest, "{}")
		return
	}

	if k == "" {
		c.JSON(http.StatusBadRequest, "{}")
		return
	}

	if _, found := s.store[g]; !found {
		c.JSON(http.StatusBadRequest, s.store[g])
	}

	c.JSON(http.StatusOK, s.store[g][k])
}
