package main

import (
	"example/middleware"
	"fmt"

	"github.com/rabbitmeow/logpaniccollector"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	lpc := logpaniccollector.New()
	lpc.LogFile = "log.log"
	lpc.AutoRemoveLog("* * * * *")
	var midd = new(middleware.Middleware)
	r.Use(midd.PanicCatcher(lpc))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello there",
		})
	})
	r.GET("/error", func(c *gin.Context) {
		listStr := []string{"a", "b", "c"}
		fmt.Println(listStr[5])
		c.JSON(500, gin.H{
			"message": "error",
		})
	})
	r.Run(":4545")
}
