package handler

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Form struct {
	File *multipart.FileHeader `form:"file"`
}

func UploadFile(c *gin.Context) {
	var form Form

	err := c.ShouldBind(&form)
	if err != nil {
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	fileName := form.File.Filename

	openedFile, err := form.File.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}

	file, err := ioutil.ReadAll(openedFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}

	err = ioutil.WriteFile("cloud/"+fileName, file, 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "err")
		return
	}

}
