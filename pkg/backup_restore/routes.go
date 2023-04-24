package backup_restore

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) Slug() string {
	return "backup"
}

func (s *Service) InitProtectedRoutes(group *gin.RouterGroup) error {
	group.GET("", s.handleCreateBackup)
	group.POST("restore", s.handleRestore)

	return nil
}

func (s *Service) handleCreateBackup(c *gin.Context) {
	zipped, zipLen, err := archiveMapToTarGz(s.BackupAll())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//if len(s.backupers) < 1 {
	//	c.JSON(http.StatusInternalServerError, "nothing to backup.")
	//	return
	//}

	timestamp := time.Now().Format("20060102-150405")
	archiveName := fmt.Sprintf("%s_xmc_backup.tar.gz", timestamp)

	c.DataFromReader(http.StatusOK, int64(zipLen), "application/gzip", zipped, map[string]string{
		"Content-Disposition": `attachment; filename="` + archiveName + `"`,
	})
}

func (s *Service) handleRestore(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error getting file: %s", err.Error()))
		return
	}

	// Open the uploaded file.
	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("Error opening file: %s", err.Error()))
		return
	}
	defer openedFile.Close()

	// Read the gzip file data.
	data, err := ioutil.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error reading file: %s", err.Error()))
		return
	}

	backup, err := unarchiveTarGzToMap(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Error unarchiving backup file: %s", err.Error()))
		return
	}

	if allErrors := s.RestoreAll(backup); len(allErrors) > 0 {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Restoring with unzipped backup data: %s", err.Error()))
	}
}
