package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/go-service-abstraction-example/pkg/app/services/abstract"
)

type ginHands struct {
	abstract.Logger
	SerName    string
	Path       string
	Latency    time.Duration
	Method     string
	StatusCode int
	ClientIP   string
	MsgStr     string
}

func Logger(serName string, l abstract.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		c.Next()
		// after request
		// latency := time.Since(t)
		// clientIP := c.ClientIP()
		// method := c.Request.Method
		// statusCode := c.Writer.Status()
		if raw != "" {
			path = path + "?" + raw
		}
		msg := c.Errors.String()
		if msg == "" {
			msg = "Request"
		}
		cData := &ginHands{
			Logger:     l,
			SerName:    serName,
			Path:       path,
			Latency:    time.Since(t),
			Method:     c.Request.Method,
			StatusCode: c.Writer.Status(),
			ClientIP:   c.ClientIP(),
			MsgStr:     msg,
		}

		logSwitch(cData)
	}
}

func logSwitch(data *ginHands) {
	var e abstract.Messager

	switch {
	case data.StatusCode >= http.StatusBadRequest && data.StatusCode < http.StatusInternalServerError:
		e = data.Warn()
	case data.StatusCode >= http.StatusInternalServerError:
		e = data.Error()
	default:
		e = data.Info()
	}

	parts := []string{
		fmt.Sprintf("method: '%v',", data.Method),
		fmt.Sprintf("path: '%v',", data.Path),
		fmt.Sprintf("resp_time: '%v',", data.Latency),
		fmt.Sprintf("status: '%v'", data.StatusCode),
	}

	e.Msgf("{%s} %s", strings.Join(parts, " "), data.MsgStr)
}
