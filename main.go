package main

import (
	"github.com/YaegerMoon/deepzoom/controller"
	"github.com/YaegerMoon/deepzoom/services"
	"github.com/gin-gonic/gin"
)

const API_PREFIX = "/api/v1"

func main() {

	router := gin.Default()
	regionDeepZoom := services.New(33, 1, 1, "png")
	controller.New(API_PREFIX, router, regionDeepZoom)
	router.Run(":8080")
}
