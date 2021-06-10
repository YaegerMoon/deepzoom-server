package controller

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/YaegerMoon/deepzoom/services"
	"github.com/gin-gonic/gin"
)

type Header struct {
	Access string `header:"Access"`
}

type Controller struct {
	prefix string
	engine *gin.Engine
	jtwkey string
}

func New(prefix string, engine *gin.Engine, jwtkey string) *Controller {

	c := &Controller{prefix, engine, jwtkey}
	c.engine.GET("/", c.healthCheck)
	group := c.engine.Group(prefix)
	{
		group.GET("/:slide/dzi", c.getDZI)
		group.GET("/:slide/dzi_files/:level/:colrow", c.getTile)
	}

	return c
}

func (controller *Controller) getDZI(c *gin.Context) {

	slidePath := c.Param("slide")
	if len(slidePath) == 0 {
		c.Status(http.StatusBadRequest)
	}
	fmt.Printf("Slide Name : %s \n", slidePath)
	services.New(slidePath, 64, 1, 1, "png", services.Area{0, 0, 0, 0})

	c.XML(http.StatusOK, gin.H{
		"width":  30,
		"height": 30,
		"level":  1,
	})
}

func (controller *Controller) getTile(c *gin.Context) {

	r, _ := regexp.Compile("([0-9]+)")
	slidePath := c.Param("slide")
	level := c.Param("level")
	colrow := r.FindAllString(c.Param("colrow"), 2)

	if len(slidePath) == 0 || len(level) == 0 || len(colrow) != 2 {
		c.Status(http.StatusBadRequest)
	}

	services.New(slidePath, 64, 1, 1, "png", services.Area{0, 0, 0, 0})
	c.JSON(http.StatusOK, gin.H{
		"file": "ImageFile",
	})
}

func (Controller *Controller) healthCheck(c *gin.Context) {
	header := Header{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}
	c.Status(http.StatusOK)
}
