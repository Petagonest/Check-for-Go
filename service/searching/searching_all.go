package search

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/logging"
)

func SearchingAll(ctx context.Context, searchStores, searchProducts, searchCategories string) ([]datastruct.Stores, []datastruct.Products, []datastruct.Categories, error) {
	db, err := logging.PembuatanKoneksi()
	var stores []datastruct.Stores
	var products []datastruct.Products
	var categories []datastruct.Categories

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM stores WHERE nama_toko LIKE '%%%s%%' OR nama_kota LIKE '%%%s%%' OR nama_kecamatan LIKE '%%%s%%' OR nama_domain LIKE '%%%s%%'",
		searchStores,
		searchStores,
		searchStores,
		searchStores)

	queryText = fmt.Sprintf("SELECT * FROM products WHERE nama_produk LIKE '%%%s%%'", searchProducts)

	queryText = fmt.Sprintf("SELECT * FROM categories WHERE nama_category LIKE '%%%s%%' OR deskripsi_category LIKE '%%%s%%'",
		searchCategories,
		searchCategories)
		
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return stores, products, categories, err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return stores, products, categories, errors.New("Maaf kata kunci yang Anda cari tidak ditemukan di database kami :(")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var store datastruct.Stores
		var product datastruct.Products
		var category datastruct.Categories

		if err = rowQuery.Scan(
			&store.Toko_id,
			&store.Nama_toko,
			&store.Kodepos_toko,
			&store.Foto_toko,
			&store.Deskripsi_toko,
			&store.Nama_domain,
			&store.Nama_kota,
			&store.Nama_kecamatan,
		); err != nil {
			return nil, nil, nil, err
		}

		stores = append(stores, store)

		if err = rowQuery.Scan(
			&product.Produk_id,
			&product.Nama_produk,
			&product.Deskripsi_produk,
			&product.Stok,
			&product.Harga_produk,
			&product.Foto_produk,
			&product.Rating_produk,
			&product.Jumlah_terjual,
			&product.Jumlah_dilihat,
			&product.Ukuran,
			&product.Warna,
		); err != nil {
			return nil, nil, nil, err
		}

		products = append(products, product)

		if err = rowQuery.Scan(
			&category.Category_id,
			&category.Nama_category,
			&category.Deskripsi_category,
		); err != nil {
			return nil, nil, nil, err
		}

		categories = append(categories, category)
	}

	return stores, products, categories, nil
}
