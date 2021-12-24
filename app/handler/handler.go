package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	err = c.SaveUploadedFile(file, "cloud/"+file.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}
}

func LoadFile(c *gin.Context) {
	filesURL, err := getAllFilesURL()
	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	c.JSON(http.StatusOK, filesURL)
}
