package main

import (
	// "login-sistem-jwt/routes/api"
	// "os"

	"login-sistem-jwt/app/Database/Migration"
	"login-sistem-jwt/routes/api"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load()
	Migration.Migrate()
	api.InitializeRoutes().Run("127.0.0.1:" + os.Getenv("PORT"))
}
