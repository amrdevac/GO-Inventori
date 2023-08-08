package api

import (
	"inventori/app/Http/Controllers/Auth/AutentikasiController"
	item "inventori/app/Http/Controllers/Item"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/register", AutentikasiController.Register)
	router.POST("/login", AutentikasiController.Login)

	router.Use(item.Constructor()).POST("/barang", item.Store)
	router.Use(item.Constructor()).GET("/barang", item.Index)
	router.Use(item.Constructor()).GET("/barang/detail", item.Detail)
	router.Use(item.Constructor()).PUT("/barang", item.Update)
	router.Use(item.Constructor()).DELETE("/barang", item.Delete)

	// authorized := router.Group("api")
	// authorized.Use(Middleware.VerifyJWT)
	// authorized.GET("/users", UserController.Index)
	// authorized.GET("/users/:email", UserController.Show)
	// authorized.PUT("/users/:email", UserController.Update)
	// authorized.DELETE("/users/:email", UserController.Destroy)
	return router
}
