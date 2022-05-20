package Migration

import (
	"login-sistem-jwt/app/Database/Mysql"
	"login-sistem-jwt/app/Http/Controllers/User/UserModel"
)

var GormConnect = Mysql.Connect()

func Migrate() {
	GormConnect.AutoMigrate(&UserModel.User{})

}
