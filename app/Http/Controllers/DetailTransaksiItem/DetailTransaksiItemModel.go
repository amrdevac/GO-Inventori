package detailtransaksiitem

import (
	"inventori/app/Database/Mysql"
	"time"
)

var GormConnect = Mysql.Connect()
var TableName = "detail_transaksi_item"
var TablePrimary = "id_detail_transaksi"

type DetailTransaksiItemTable struct {
	IdDetailTransaksi int       `json:"id_detail_transaksi" gorm:"column:id_detail_transaksi;type:integer;not null;primaryKey;autoIncrement:true"`
	RefIdBarang       int       `json:"ref_id_barang" gorm:"column:ref_id_barang; type:integer; not null;"`
	RefIdTransaksi    int       `json:"ref_id_transaksi" gorm:"column:ref_id_transaksi; type:integer; not null;"`
	Qty               int       `json:"qty" gorm:"column:qty; type:integer; not null;"`
	CreatedAt         time.Time `json:"created_at" gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP()"`
}

func (DetailTransaksiItemTable) TableName() string {
	return TableName
}
