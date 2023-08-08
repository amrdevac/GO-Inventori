package Migration

import (
	"inventori/app/Database/Mysql"
	item "inventori/app/Http/Controllers/Item"
	"inventori/app/Http/Controllers/User/UserModel"
)

var GormConnect = Mysql.Connect()

func Migrate() {
	GormConnect.AutoMigrate(&item.ItemTable{})
	GormConnect.AutoMigrate(&UserModel.User{})
}
