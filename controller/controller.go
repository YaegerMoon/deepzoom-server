package controller

import (
	"fmt"
	"net/http"
	"regexp"

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
	// area := services.Area{Top: 0, Left: 0, Right: 0, Bottom: 0}
	// services.New(slidePath, 64, 1, 1, "png", area)

	c.XML(http.StatusOK, gin.H{
		"width":  30,
		"height": 30,
		"level":  1,
	})
}

func (ctl *Controller) getTile(c *gin.Context) {

	r, _ := regexp.Compile("([0-9]+)")
	slidePath := c.Param("slide")
	level := c.Param("level")
	colrow := r.FindAllString(c.Param("colrow"), 2)

	if len(slidePath) == 0 || len(level) == 0 || len(colrow) != 2 {
		c.Status(http.StatusBadRequest)
	}
	// area := services.Area{Top: 0, Left: 0, Right: 0, Bottom: 0}
	// services.New(slidePath, 64, 1, 1, "png", area)
	c.JSON(http.StatusOK, gin.H{
		"file": "ImageFile",
	})
}

func (ctl *Controller) healthCheck(c *gin.Context) {
	h := Header{}
	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"s": 4321})
}

func (ctl *Controller) Run(port string) {
	ctl.engine.Run(port)
}
