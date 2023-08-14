package controllers

import (
	"inventori/app/Database/Mysql"

	"github.com/gin-gonic/gin"
)

var MaindGormConnect = Mysql.Connect()
var GlobalGContext *gin.Context

func MainConstructor() gin.HandlerFunc {
	return func(c *gin.Context) {
		GlobalGContext = c
		c.Next()
	}
}

