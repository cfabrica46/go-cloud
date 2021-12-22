package main

import (
	"net/http"

	"github.com/cfabrica46/go-cloud/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))
}

func setupRouter() (r *gin.Engine) {
	r = gin.Default()

	setCors(r)

	s := r.Group("/api/v1")
	s.POST("/upload", handler.UploadFile)
	s.GET("/load", handler.LoadFile)
	return
}

func runRedirectHTTPToHTTPS(portHTTP, portHTTPS string) {
	httpRouter := gin.Default()

	setCors(httpRouter)

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
