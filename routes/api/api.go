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

	authorized := router
	// authorized.Use(Middleware.VerifyJWT)

	authorized.Use(item.Constructor()).POST("/barang", item.Store)
	authorized.Use(item.Constructor()).GET("/barang", item.Index)
	authorized.Use(item.Constructor()).GET("/barang/detail", item.Detail)
	authorized.Use(item.Constructor()).PUT("/barang", item.Update)
	authorized.Use(item.Constructor()).DELETE("/barang", item.Delete)

	authorized.Use(transaksiitem.Constructor()).POST("/transaksi-item", transaksiitem.Store)
	authorized.Use(transaksiitem.Constructor()).GET("/transaksi-item", transaksiitem.Index)
	authorized.Use(transaksiitem.Constructor()).GET("/transaksi-item/detail", transaksiitem.Detail)
	authorized.Use(transaksiitem.Constructor()).GET("/transaksi-item/detail-full", transaksiitem.DetailFull)
	authorized.Use(transaksiitem.Constructor()).DELETE("/transaksi-item", transaksiitem.Delete)
	
	
	// authorized := router.Group("api")
	// authorized.Use(Middleware.VerifyJWT)
	// authorized.GET("/users", UserController.Index)
	// authorized.GET("/users/:email", UserController.Show)
	// authorized.PUT("/users/:email", UserController.Update)
	// authorized.DELETE("/users/:email", UserController.Destroy)
	return router
}
