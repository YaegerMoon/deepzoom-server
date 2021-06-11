package main

import (
	"github.com/YaegerMoon/deepzoom/controller"
	"github.com/gin-gonic/gin"
)

const API_PREFIX = "/api/v1"

//SHOULD BE REMOVED ON PRODUCTION
const JWT_KEY = "hello"

func main() {

	c := controller.New(API_PREFIX, gin.Default(), JWT_KEY)
	c.Run(":8080")
}
