package Migration

import (
	"inventori/app/Database/Mysql"
	detailtransaksiitem "inventori/app/Http/Controllers/DetailTransaksiItem"
	item "inventori/app/Http/Controllers/Item"
	transaksiitem "inventori/app/Http/Controllers/TransaksiItem"
	"inventori/app/Http/Controllers/User/UserModel"
)

var GormConnect = Mysql.Connect()

func Migrate() {
	GormConnect.AutoMigrate(&item.ItemTable{})
	GormConnect.AutoMigrate(&UserModel.User{})
	GormConnect.AutoMigrate(&transaksiitem.TransaksiItemTable{})
	GormConnect.AutoMigrate(&detailtransaksiitem.DetailTransaksiItemTable{})
}
