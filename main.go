package main

import (
	"github.com/YaegerMoon/deepzoom/controller"
	"github.com/gin-gonic/gin"
)

const API_PREFIX = "/api/v1"

func main() {

	router := gin.Default()

	controller.New(API_PREFIX, router)
	router.Run(":8080")
}
