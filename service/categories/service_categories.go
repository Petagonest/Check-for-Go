package categories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/logging"
)

const (
	table          = "categories"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll categories
func GetAll(ctx context.Context) ([]datastruct.Categories, error) {

	var categories []datastruct.Categories

	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By nama_category ASC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category datastruct.Categories

		if err = rowQuery.Scan(
			&category.Category_id,
			&category.Nama_category,
			&category.Deskripsi_category); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func SearchingCategories(ctx context.Context, searchCategories string) ([]datastruct.Categories, error) {
	db, err := logging.PembuatanKoneksi()
	var categories []datastruct.Categories

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}
	queryText := fmt.Sprintf("SELECT * FROM categories WHERE nama_category LIKE '%%%s%%' OR deskripsi_category LIKE '%%%s%%'",
		searchCategories,
		searchCategories)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return categories, err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return categories, errors.New("Maaf kata kunci yang Anda cari tidak ditemukan di database kami :(")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	rowQuery, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category datastruct.Categories

		if err = rowQuery.Scan(
			&category.Category_id,
			&category.Nama_category,
			&category.Deskripsi_category,
		); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

// Insert categories
func Insert(ctx context.Context, category datastruct.Categories) error {
	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah ID yang dicari gaada :(", err)
	}

	queryText := fmt.Sprintf("INSERT INTO categories (category_id, nama_category, deskripsi_category) VALUES (nextval('category_id'),'%v','%v')",
		category.Nama_category,
		category.Deskripsi_category,
	)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

// Update categories
func Update(ctx context.Context, category datastruct.Categories, id string) error {

	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah ID yang dicari gaada :(", err)
	}

	queryText := fmt.Sprintf("UPDATE %v set nama_category ='%s', deskripsi_category ='%s' where category_id = %s",
		table,
		category.Nama_category,
		category.Deskripsi_category,
		id,
	)
	fmt.Println(queryText)

	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil
}

// Delete categories
func Delete(ctx context.Context, id string) error {
	db, err := logging.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah ID yang dicari gaada :(", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where category_id = %s", table, id)

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
