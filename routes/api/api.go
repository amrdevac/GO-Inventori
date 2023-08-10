package api

import (
	controllers "inventori/app/Http/Controllers"
	"inventori/app/Http/Controllers/Auth/AutentikasiController"
	item "inventori/app/Http/Controllers/Item"
	transaksiitem "inventori/app/Http/Controllers/TransaksiItem"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(controllers.MainConstructor())
	router.POST("/register", AutentikasiController.Register)
	router.POST("/login", AutentikasiController.Login)

	router.Use(item.Constructor()).POST("/barang", item.Store)
	router.Use(item.Constructor()).GET("/barang", item.Index)
	router.Use(item.Constructor()).GET("/barang/detail", item.Detail)
	router.Use(item.Constructor()).PUT("/barang", item.Update)
	router.Use(item.Constructor()).DELETE("/barang", item.Delete)

	router.Use(transaksiitem.Constructor()).POST("/transaksi-item", transaksiitem.Store)
	router.Use(transaksiitem.Constructor()).GET("/transaksi-item", transaksiitem.Index)
	router.Use(transaksiitem.Constructor()).GET("/transaksi-item/detail", transaksiitem.Detail)
	router.Use(transaksiitem.Constructor()).DELETE("/transaksi-item", transaksiitem.Delete)
	
	
	// authorized := router.Group("api")
	// authorized.Use(Middleware.VerifyJWT)
	// authorized.GET("/users", UserController.Index)
	// authorized.GET("/users/:email", UserController.Show)
	// authorized.PUT("/users/:email", UserController.Update)
	// authorized.DELETE("/users/:email", UserController.Destroy)
	return router
}
