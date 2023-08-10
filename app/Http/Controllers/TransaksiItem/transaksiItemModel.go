package transaksiitem

import (
	"inventori/app/Database/Mysql"
	"time"
)

var GormConnect = Mysql.Connect()
var TableName = "transaksi_item"
var TablePrimary = "id_transaksi"


type TransaksiItemTable struct {
	IdTransaksi   int       `json:"id_transaksi" gorm:"column:id_transaksi;type:integer;not null;primaryKey;autoIncrement:true"`
	TipeTransaksi string    `json:"tipe_transaksi" gorm:"column:tipe_transaksi;type:enum('MASUK', 'KELUAR');not null;default:MASUK"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP()"`
}

func (TransaksiItemTable) TableName() string {
	return TableName
}
