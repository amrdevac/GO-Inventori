package detailtransaksiitem

import (
	"github.com/gin-gonic/gin"
)

var ArrTipeTransaksi []string

func Constructor() gin.HandlerFunc {
	return func(c *gin.Context) {
		ArrTipeTransaksi = []string{"MASUK", "KELUAR"}
		privateGContext = c
		c.Next()
	}
}
