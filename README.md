# What is it ?

The [Golang](https://golang.org) package for writing log or http panic into file. You can combine this package with [Filebeat](https://www.elastic.co/products/beats/filebeat) or [Promtail](https://github.com/grafana/loki/tree/master/docs/clients/promtail).

## Install

`go get -u github.com/rabbitmeow/logpaniccollector`

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

//Middleware ...
type Middleware struct{}

// PanicCatcher is use for collecting panic that happened in endpoint
func (w *Middleware) PanicCatcher() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			isPanic := logpaniccollector.RecoverPanic(c.Request.RequestURI, recover())
			if isPanic {
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

# To Do

- [x] simplify middleware
- [ ] feature cron for auto clean the file
- [ ] feature enable/disable [the Filebeat multiline support](https://www.elastic.co/guide/en/beats/filebeat/current/multiline-examples.html) (currently is enabled)
- [ ] feature enable/disable panic log 1 line mode as a follow up from [Promtail issue](https://github.com/grafana/loki/issues/74)
