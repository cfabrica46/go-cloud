package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func setupRouter() (r *gin.Engine) {
	r = gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./statics", true)))

	s := r.Group("/api/v1")
	// s.POST("/upload", uploadFile)
	return
}
