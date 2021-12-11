package products

import (
	"Check-for-Go/config"
	"Check-for-Go/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table          = "products"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll products
func GetAll(ctx context.Context) ([]models.Products, error) {

	var products []models.Products

	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By produk_id ASC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var produk models.Products

		if err = rowQuery.Scan(
			&produk.Produk_id,
			&produk.Nama_produk,
			&produk.Deskripsi_produk,
			&produk.Stok,
			&produk.Harga_produk,
			&produk.Foto_produk,
			&produk.Rating_produk,
			&produk.Jumlah_terjual,
			&produk.Jumlah_dilihat,
			&produk.Ukuran,
			&produk.Warna); err != nil {
			return nil, err
		}

		products = append(products, produk)
	}

	return products, nil
}

// Insert products
func Insert(ctx context.Context, produk models.Products) error {
	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (produk_id, nama_produk, deskripsi_produk, stok, harga_produk, foto_produk, rating_produk, jumlah_terjual, jumlah_dilihat, ukuran, warna) values('%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v')", table,
		produk.Produk_id,
		produk.Nama_produk,
		produk.Deskripsi_produk,
		produk.Stok,
		produk.Harga_produk,
		produk.Foto_produk,
		produk.Rating_produk,
		produk.Jumlah_terjual,
		produk.Jumlah_dilihat,
		produk.Ukuran,
		produk.Warna,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update products
func Update(ctx context.Context, produk models.Products, id string) error {

	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama_produk ='%s', deskripsi_produk ='%s', stok ='%d', harga_produk ='%d', foto_produk ='%s', rating_produk ='%f', jumlah_terjual ='%d', jumlah_dilihat ='%d', ukuran ='%s', warna ='%s' where produk_id = %s",
		table,
		produk.Nama_produk,
		produk.Deskripsi_produk,
		produk.Stok,
		produk.Harga_produk,
		produk.Foto_produk,
		produk.Rating_produk,
		produk.Jumlah_terjual,
		produk.Jumlah_dilihat,
		produk.Ukuran,
		produk.Warna,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete products
func Delete(ctx context.Context, id string) error {
	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where produk_id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("yah ID yang dicari gaada :(")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
