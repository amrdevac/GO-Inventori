package item

import (
	"inventori/app/Database/Mysql"
	"time"
)

var GormConnect = Mysql.Connect()

type ItemTable struct {
	IdBarang  int       `form:"id_barang" json:"id_barang" gorm:"column:id_barang;type:integer;not null;primaryKey;autoIncrement:true"`
	Nama      string    `form:"nama" json:"nama" binding:"required" gorm:"column:nama;type:varchar(100);not null"`
	Satuan    string    `form:"satuan" json:"satuan" binding:"required" gorm:"column:satuan;type:varchar(100);not null"`
	Harga     int       `form:"harga" json:"harga" binding:"required" gorm:"column:harga;type:varchar(100);not null"`
	Stok      string    `form:"stok" json:"stok" binding:"required" gorm:"column:stok;type:varchar(100);not null"`
	CreatedAt time.Time `form:"created_at" json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP()"`
}

func (ItemTable) TableName() string {
	return "barang"
}

// func (ListBarangRequest) TableName() string {
// 	return "barang"
// }
