package item

import (
	"github.com/gin-gonic/gin"
)

func Constructor() gin.HandlerFunc {
	return func(c *gin.Context) {
		privateGContext = c
		c.Next()
	}
}
