package middlewares

import (
	"github.com/gin-gonic/gin"
)

type Cors struct {
	//
}

func (cs *Cors) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
