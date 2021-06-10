package main

import (
	"github.com/YaegerMoon/deepzoom/controller"
	"github.com/gin-gonic/gin"
)

const API_PREFIX = "/api/v1"

//SHOULD BE REMOVED ON PRODUCTION
const JWT_KEY = "hello"

func main() {

	router := gin.Default()
	controller.New(API_PREFIX, router)
	router.Run(":8080")
}
