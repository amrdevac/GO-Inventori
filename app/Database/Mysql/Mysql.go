package Mysql

import (
	"fmt"
	ErrorHandler "inventori/app/Provider/ErrorHandler"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	loadEnvErrorCheck := godotenv.Load()
	ErrorHandler.Err(loadEnvErrorCheck).Check(".env").Fatal()

	mysqlLocalConnect := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=true",
		os.Getenv("databaseUser"),
		os.Getenv("databasePassword"),
		os.Getenv("databaseName"),
		// os.Getenv("mysql"),
		// os.Getenv("localhost"),
		// os.Getenv("3306"),
	)

	fmt.Println(mysqlLocalConnect)
	// database, connectDbErrorCheck := gorm.Open("mysql", mysqlLocalConnect)

	
	database, connectDbErrorCheck := gorm.Open(mysql.Open(mysqlLocalConnect))
	ErrorHandler.Err(connectDbErrorCheck).Check("Connect Mysql").Fatal()

	return database
}
