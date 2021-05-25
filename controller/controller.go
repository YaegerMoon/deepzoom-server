package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	prefix string
	engine *gin.Engine
}

func New(prefix string, engine *gin.Engine) *Controller {

	group := engine.Group(prefix)
	{
		group.GET("/:slide/dzi", getDZI)
		group.GET("/:slide/dzi_files/:level/:colrow", getTile)
	}

	c := &Controller{prefix, engine}

	return c
}

func getDZI(c *gin.Context) {

	slideName := c.Param("slide")

	fmt.Printf("Slide Name : %s \n", slideName)

	c.XML(http.StatusOK, gin.H{
		"width":  30,
		"height": 30,
		"level":  1,
	})
}

func getTile(c *gin.Context) {

	slideName := c.Param("slide")
	level := c.Param("level")
	colrow := c.Param("colrow")

	fmt.Printf("Slide Name : %s \n", slideName)
	fmt.Printf("Level : %s \n", level)
	fmt.Printf("Col Row : %s \n", colrow)

	c.JSON(http.StatusOK, gin.H{
		"file": "ImageFile",
	})
}
