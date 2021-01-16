package handler

import "github.com/gin-gonic/gin"

func Getting(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
