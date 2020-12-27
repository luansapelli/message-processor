package controllers

import (
	"github.com/gin-gonic/gin"
)

func Router () {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	api := router.Group("/golang/event-processor/")
	api.GET("/health-check", HealthCheck)

	router.Run(":80")
}

