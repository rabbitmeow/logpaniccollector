package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmeow/logpaniccollector"
)

//Middleware is
type Middleware struct{}

// PanicCatcher is use for collecting panic that happened in endpoint
func (w *Middleware) PanicCatcher(lpc *logpaniccollector.LogPanic) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			isPanic := lpc.RecoverPanic(c.Request.RequestURI, recover())
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
