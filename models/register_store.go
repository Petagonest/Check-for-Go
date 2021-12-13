package models

type (
	//Stores
	Stores struct {
		Toko_id        int64  `json:"toko_id"`
		Nama_toko      string `json:"nama_toko"`
		Kodepos_toko   string `json:"kodepos_toko"`
		Nama_kota      string `json:"nama_kota"`
		Nama_kecamatan string `json:"nama_kecamatan"`
		Foto_toko      string `json:"foto_toko"`
		Deskripsi_toko string `json:"deskripsi_toko"`
		Nama_domain    string `json:"nama_domain"`
	}

	//Products
	Products struct {
		Produk_id        int64   `json:"produk_id"`
		Nama_produk      string  `json:"nama_produk"`
		Deskripsi_produk string  `json:"deskripsi_produk"`
		Stok             int64   `json:"stok"`
		Harga_produk     int64   `json:"harga_produk"`
		Foto_produk      string  `json:"foto_produk"`
		Rating_produk    float64 `json:"rating_produk"`
		Jumlah_terjual   int64   `json:"jumlah_terjual"`
		Jumlah_dilihat   int64   `json:"jumlah_dilihat"`
		Ukuran           string  `json:"ukuran"`
		Warna            string  `json:"warna"`
		Toko_id          int64   `json:"toko_id"`
	}

	//Categories
	Categories struct {
		Category_id        int64  `json:"category_id"`
		Nama_category      string `json:"nama_category"`
		Deskripsi_category string `json:"deskripsi_category"`
		Produk_id          int64  `json:"produk_id"`
	}
)
