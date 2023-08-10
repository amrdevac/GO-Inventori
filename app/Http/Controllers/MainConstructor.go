package controllers

import (
	"inventori/app/Database/Mysql"

	"github.com/gin-gonic/gin"
)

var MaindGormConnect = Mysql.Connect()
func MainConstructor() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}

