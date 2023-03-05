package static_assets

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (s *Service) staticWebUIHandler(c *gin.Context) {
	_, file := path.Split(c.Request.RequestURI)
	ext := filepath.Ext(file)

	if file == "" || ext == "" {
		file = "index.html"
	}

	file = fmt.Sprintf("embedded/%s", file)

	data, readErr := s.ReadFile(file)
	if readErr != nil {
		c.Redirect(http.StatusPermanentRedirect, "404.html")
		c.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
		return
	}

	contentTypeToUse := getMIMEFromFileExtension(ext)
	if contentTypeToUse == "" {
		contentTypeToUse = http.DetectContentType(data)
	}

	c.Data(http.StatusOK, contentTypeToUse, data)

	return
}

func getMIMEFromFileExtension(ext string) (result string) {
	return map[string]string{
		".js":  "text/javascript",
		".css": "text/css",
	}[ext]
}
