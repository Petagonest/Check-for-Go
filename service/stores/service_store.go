package stores

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Petagonest/Check-for-Go/logging"
	"github.com/Petagonest/Check-for-Go/datastruct"
)

const (
	table          = "provinsi"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll stores
func GetAll(ctx context.Context) ([]datastruct.Provinsi, error) {

	var provinsi []datastruct.Provinsi

	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	// queryText := fmt.Sprintf("SELECT * FROM %v Order By toko_id ASC", table)
	// queryText := fmt.Sprintf("SELECT * FROM %v where toko_id = 3", table)
	queryText := fmt.Sprintf("SELECT * FROM %v Order By id ASC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var prov datastruct.Provinsi

		if err = rowQuery.Scan(
			&prov.Id,
			&prov.Name,
			// &store.Kodepos_toko,
			// &store.Foto_toko,
			// &store.Deskripsi_toko,
			// &store.Nama_domain,
			// &store.Nama_kota,
			// &store.Nama_kecamatan
			); err != nil {
			return nil, err
		}
		if err = rowQuery.Scan(
			&prov.Id,
			&prov.Name); err != nil {
			return nil, err
		}

		provinsi = append(provinsi, prov)
	}

	return provinsi, nil
}

// Insert stores
func Insert(ctx context.Context, store datastruct.Stores) error {
	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	checkUsername := fmt.Sprintf("SELECT FROM %v where nama_toko = %s", table, store.Nama_toko)
	if checkUsername != nil {
		log.Fatal("Nama toko sudah ada :(", err)
	} 

	queryText := fmt.Sprintf("INSERT INTO %v (toko_id, nama_toko, kodepos_toko, foto_toko, deskripsi_toko, nama_domain, nama_kota, nama_kecamatan) VALUES ('%v','%v','%v','%v','%v','%v','%v','%v')", table,
		store.Toko_id,
		store.Nama_toko,
		store.Kodepos_toko,
		store.Foto_toko,
		store.Deskripsi_toko,
		store.Nama_domain,
		store.Nama_kota,
		store.Nama_kecamatan,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update stores
func Update(ctx context.Context, store datastruct.Stores, id string) error {

	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama_toko ='%s', kodepos_toko ='%s', foto_toko ='%s', deskripsi_toko ='%s', nama_domain ='%s', nama_kota ='%s', nama_kecamatan ='%s' where toko_id = %s",
		table,
		store.Nama_toko,
		store.Kodepos_toko,
		store.Foto_toko,
		store.Deskripsi_toko,
		store.Nama_domain,
		store.Nama_kota,
		store.Nama_kecamatan,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete stores
func Delete(ctx context.Context, id string) error {
	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where toko_id = %s", table, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("Yah ID yang dicari gaada :(")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
