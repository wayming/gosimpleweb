package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wayming/gosimpleweb/src/handler"
)

func main() {
	router := gin.Default()
	router.GET("/ping", handler.Getting)
	router.Run() // listen and serve on 0.0.0.0:8080
}
