package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "err")
		return
	}

	err = c.SaveUploadedFile(file, "cloud/"+file.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "err")
		return
	}

}

func LoadFile(c *gin.Context) {
	log.Println("uwu")

	c.File("cloud/beta.png")

	/* _, err := c.Writer.Write([]byte("hola"))
	if err != nil {
		log.Println(err)
	} */

	c.JSON(http.StatusOK, "1")

}
