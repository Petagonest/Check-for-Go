package categories

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
	table          = "categories"
	layoutDateTime = "2021-09-27 03:05:05"
)

// GetAll categories
func GetAll(ctx context.Context) ([]models.Categories, error) {

	var categories []models.Categories

	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah gagal connect ke Postgress :(", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By category_id ASC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var category models.Categories

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

// Insert categories
func Insert(ctx context.Context, category models.Categories) error {
	db, err := config.PembuatanKoneksi()

	if err != nil {
		log.Fatal("Yah ID yang dicari gaada :(", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (category_id, nama_category, deskripsi_category) values('%v','%v','%v')", table,
		category.Category_id,
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
func Update(ctx context.Context, category models.Categories, id string) error {

	db, err := config.PembuatanKoneksi()

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
	db, err := config.PembuatanKoneksi()

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
