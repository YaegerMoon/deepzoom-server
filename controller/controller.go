package controller

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/YaegerMoon/deepzoom/services"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	prefix   string
	engine   *gin.Engine
	deepzoom *services.RegionDeepZoom
}

func New(prefix string, engine *gin.Engine, deepzoom *services.RegionDeepZoom) *Controller {

	c := &Controller{prefix, engine, deepzoom}

	group := c.engine.Group(prefix)
	{
		group.GET("/:slide/dzi", c.getDZI)
		group.GET("/:slide/dzi_files/:level/:colrow", c.getTile)
	}

	return c
}

func (controller *Controller) getDZI(c *gin.Context) {

	slide := c.Param("slide")
	if len(slide) == 0 {
		c.Status(http.StatusBadRequest)
	}
	fmt.Printf("Slide Name : %s \n", slide)

	c.XML(http.StatusOK, gin.H{
		"width":  30,
		"height": 30,
		"level":  1,
	})
}

func (controller *Controller) getTile(c *gin.Context) {

	r, _ := regexp.Compile("([0-9]+)")
	slide := c.Param("slide")
	level := c.Param("level")
	colrow := r.FindAllString(c.Param("colrow"), 2)

	if len(slide) == 0 || len(level) == 0 || len(colrow) != 2 {
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, gin.H{
		"file": "ImageFile",
	})
}
