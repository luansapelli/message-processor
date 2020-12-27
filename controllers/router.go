package controllers

import "github.com/gin-gonic/gin"

func Router() {

	router := gin.Default()

	api := router.Group("/golang/event-processor/")
	api.GET("/health-check", healthCheck)

	router.Run(":80")
}
