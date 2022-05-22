package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func metadata(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func index(c *gin.Context) {
	c.String(200, "ok")
}

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/metadata", metadata)
	r.GET("/", index)
	return r
}

func Run() {
	r := SetupRouter()
	log.Println("running on 8080")
	r.Run(":8080")
}
