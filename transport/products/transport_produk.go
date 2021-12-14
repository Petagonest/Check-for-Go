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
func GetProducts(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	prd, err := products.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	logging.ResponseJSON(w, prd, http.StatusOK)
}

// Create
// PostProducts
func PostProducts(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	if r.Header.Get("Content-Type") != "aCheck-for-Goication/json" {
		http.Error(w, "Gunakan content type aCheck-for-Goication / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var prd datastruct.Products

	if err := json.NewDecoder(r.Body).Decode(&prd); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := products.Insert(ctx, prd); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateProducts
func UpdateProducts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "aCheck-for-Goication/json" {
		http.Error(w, "Gunakan content type aCheck-for-Goication / json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var prd datastruct.Products

	if err := json.NewDecoder(r.Body).Decode(&prd); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idProducts = ps.ByName("id")

	if err := products.Update(ctx, prd, idProducts); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusCreated)
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

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////