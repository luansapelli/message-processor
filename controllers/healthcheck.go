package controllers

import "github.com/gin-gonic/gin"

func healthCheck(context *gin.Context) {

	context.JSON(200, gin.H{
		"status": "alive and kicking!",
	})
}
