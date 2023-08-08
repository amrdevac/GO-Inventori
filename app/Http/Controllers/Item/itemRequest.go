package item

type ListBarangRequest struct {
	Limit   int    `json:"limit"`
	Offset  int    `json:"offset"`
	F_nama  string `json:"f_nama"`
	F_harga int    `json:"f_harga"`
}

type DetailItemRequest struct {
	IdBarang int `form:"id_barang" json:"id_barang"`
}

type UpdateItemRequest struct {
	IdBarang  int       `form:"id_barang" json:"id_barang" binding:"required"`
	Nama      string    `form:"nama" json:"nama" binding:"required" `
	Satuan    string    `form:"satuan" json:"satuan" binding:"required" `
	Harga     int       `form:"harga" json:"harga" binding:"required" `
	Stok      string    `form:"stok" json:"stok" binding:"required" `
}
