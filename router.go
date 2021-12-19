package main

import (
	"net/http"

	"github.com/cfabrica46/go-cloud/handler"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "hola")
}

func setupRouter() (r *gin.Engine) {
	r = gin.Default()
	r.GET("/", test)
	s := r.Group("/api/v1")
	s.POST("/upload", handler.UploadFile)
	return
}

func runRedirectHTTPToHTTPS(portHTTP, portHTTPS string) {
	httpRouter := gin.Default()
	httpRouter.Any("/*path", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "https://localhost:"+portHTTPS+c.Request.RequestURI)
	})

	go httpRouter.Run(":" + portHTTP)
}

func runServer(portHTTP, portHTTPS string) (err error) {
	runRedirectHTTPToHTTPS(portHTTP, portHTTPS)

	r := setupRouter()
	err = r.RunTLS(":"+portHTTPS, "server.crt", "server.key")
	if err != nil {
		return
	}
	return
}
