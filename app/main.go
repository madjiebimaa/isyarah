package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/madjiebimaa/isyarah/app/config"
)

func main() {
	config.GetEnv()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.POST("/api/locations", func(c *gin.Context) {})

	if err := r.Run(":3000"); err != nil {
		log.Fatal(err)
	}
}
