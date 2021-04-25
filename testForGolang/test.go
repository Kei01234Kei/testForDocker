package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	t := time.Now()
	engine:= gin.Default()
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", gin.H{
			"message": "hello world",
                        "time": t,
		})
	})
	engine.GET("/test", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html", gin.H{
			"message": "TestPage",
                        "time": t,
		})
	})
	engine.Run(":3000")
}
