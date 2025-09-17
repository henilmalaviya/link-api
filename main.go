package main

import (
	"fmt"
	"log"

	"api.link.henil.dev/docs"
	"api.link.henil.dev/internal/env"
	"api.link.henil.dev/server/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	r := gin.New()

	// Custom logger middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	routes.Register(r)

	r.GET("/health", func(c *gin.Context) {
		c.String(200, "OK")
	})

	port := fmt.Sprintf(":%d", env.Get().Port)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
