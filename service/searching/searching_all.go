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

func SearchingAll(ctx context.Context, searchAll string) ([]datastruct.Stores, []datastruct.Products, []datastruct.Categories, error) {
	db, err := logging.PembuatanKoneksi()
	var searchAll []datastruct.Stores, []datastruct.Products, []datastruct.Categories

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM stores WHERE nama_toko LIKE '%%%s%%' OR nama_kota LIKE '%%%s%%' OR nama_kecamatan LIKE '%%%s%%' OR nama_domain LIKE '%%%s%%'",
		searchAll,
		searchAll,
		searchAll,
		searchAll)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return stores, err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return stores, errors.New("Maaf kata kunci yang Anda cari tidak ditemukan di database kami :(")
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
			return nil, err
		}

		stores = append(stores, store)
	}

	return stores, nil
}
