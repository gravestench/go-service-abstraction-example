package static_assets

import (
	"net/http"
	"path"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) staticWebUIHandler(c *gin.Context) {
	file := c.Request.RequestURI

	_, baseFilename := path.Split(file)
	ext := filepath.Ext(baseFilename)

	if file == "" || ext == "" {
		file = "index.html"
	}

	data, readErr := m.ReadFile(file)
	if readErr != nil {
		m.log.Warn().Msgf("[%s] reading file: %v", m.Name(), readErr)
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
