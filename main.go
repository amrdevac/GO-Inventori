package main

import (
	"inventori/app/Database/Migration"
	"inventori/routes/api"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	Migration.Migrate()
	api.InitializeRoutes().Run("127.0.0.1:" + os.Getenv("PORT"))
}
