package Mysql

import (
	ErrorHandler "login-sistem-jwt/app/Provider/ErrorHandler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func Connect()*gorm.DB {

	loadEnvErrorCheck := godotenv.Load()
	ErrorHandler.Err(loadEnvErrorCheck).Check(".env").Fatal()

	database, connectDbErrorCheck := gorm.Open("mysql", "root:@/golang_tes_login_jwt?charset=utf8&parseTime=true")
	ErrorHandler.Err(connectDbErrorCheck).Check("Connect Mysql").Fatal()
	

	return database
}
