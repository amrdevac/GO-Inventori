package api

import (
	"login-sistem-jwt/app/Http/Controllers/Auth/AutentikasiController"
	"login-sistem-jwt/app/Http/Controllers/Auth/Middleware"
	"login-sistem-jwt/app/Http/Controllers/User/UserController"

	"github.com/gin-gonic/gin"
)


func InitializeRoutes() *gin.Engine {
	router := gin.Default()
	
	router.POST("/register", AutentikasiController.Register)
	router.POST("/login", AutentikasiController.Login)

	authorized := router.Group("api")
	authorized.Use(Middleware.VerifyJWT)
	authorized.GET("/users", UserController.Index)
	authorized.GET("/users/:email", UserController.Show)
	authorized.PUT("/users/:email", UserController.Update)
	authorized.DELETE("/users/:email", UserController.Destroy)
	return router
}
