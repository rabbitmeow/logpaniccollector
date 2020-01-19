# What is it ?

The [Golang](https://golang.org) package for writing log or http panic into file. You can combine this package with [Filebeat](https://www.elastic.co/products/beats/filebeat) or [Promtail](https://github.com/grafana/loki/tree/master/docs/clients/promtail).

## Install

`go get github.com/rabbitmeow/logpaniccollector`

You can also use vendoring tools like [Govendor](https://github.com/kardianos/govendor), [Dep](https://github.com/golang/dep), or something else.

## Docs

<https://godoc.org/github.com/rabbitmeow/logpaniccollector>

## Usage

Please make your middleware to use the [WritePanic](https://godoc.org/github.com/rabbitmeow/logpaniccollector#WritePanic) function. Below is the example of Gin middleware

middleware/middleware.go :
```
package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/rabbitmeow/logpaniccollector"
)

//Middleware is
type Middleware struct{}

// PanicCatcher is use for collecting panic that happened in endpoint
func (w *Middleware) PanicCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			rec := recover()
			if rec != nil {
				errPanic := fmt.Sprintf("Endpoint: %s - panic: %v", c.Request.RequestURI, rec)
				logpaniccollector.WritePanic(errPanic, debug.Stack())
				c.JSON(500, gin.H{
					"status":  500,
					"message": "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
```

main.go
```
package main

import (
	"[YOUR_MODULE_NAME]/middleware"
	"github.com/rabbitmeow/logpaniccollector"
	"github.com/gin-gonic/gin"
)

func main() {
    //...
	r := gin.Default()
	logpaniccollector.SetFile("server.log")
	var midd = new(middleware.Middleware)
	r.Use(midd.PanicCatcher())
	logpaniccollector.WriteLog("test log")
    //...
}
```