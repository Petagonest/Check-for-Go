package transport_products

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/logging"
	"github.com/Petagonest/Check-for-Go/service/products"
	"github.com/julienschmidt/httprouter"
)
//--------Products----------//
// Read
// GetProducts
func GetAllProducts(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	product, err := products.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	logging.ResponseJSON(w, product, http.StatusOK)
}

func SearchProducts(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var searchProducts = ps.ByName("search")
	products, err := products.SearchingProducts(ctx, searchProducts)

	if err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	logging.ResponseJSON(w, products, http.StatusOK)
}

// Create
// PostProducts
func PostProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var product datastruct.Products

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := products.Insert(ctx, product); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)

}

// UpdateProducts
func UpdateProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var product datastruct.Products

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idProducts = ps.ByName("id")

	if err := products.Update(ctx, product, idProducts); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)
}

// Delete
// DeleteProducts
func DeleteProducts(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idProducts = ps.ByName("id")

	if err := products.Delete(ctx, idProducts); err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////