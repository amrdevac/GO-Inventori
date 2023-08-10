package detailtransaksiitem

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
	RefIdBarang       int       `form:"ref_id_barang" json:"ref_id_barang" binding:"requierd"`
	RefIdTransaksi    int       `form:"ref_id_transaksi" json:"ref_id_transaksi" binding:"requierd"`
	Qty               int       `form:"qty" json:"qty" binding:"requierd"`
}
