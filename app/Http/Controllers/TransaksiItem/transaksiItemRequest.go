package transaksiitem

import (
	detailtransaksiitem "inventori/app/Http/Controllers/DetailTransaksiItem"
	"time"
)

type ListRequest struct {
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	F_nama  string `json:"f_nama"`
	F_harga int    `json:"f_harga"`
}

type DetailRequest struct {
	Id_Transaksi int `form:"id_transaksi" json:"id_transaksi" binding:"required"`
}

type InsertRequest struct {
	Id_Transaksi   int                                 `form:"id_transaksi" json:"id_transaksi"`
	Tipe_Transaksi string                              `form:"tipe_transaksi" json:"tipe_transaksi" binding:"required"`
	RefBarang      []detailtransaksiitem.InsertRequest `form:"ref_barang" json:"ref_barang" binding:"required"`
	CreatedAt      time.Time                           `form:"created_at" json:"created_at" `
}


