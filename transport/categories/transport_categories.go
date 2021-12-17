package transport_categories

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/logging"
	"github.com/Petagonest/Check-for-Go/service/categories"
	"github.com/julienschmidt/httprouter"
)

//--------Categories----------//
// Read
// GetCategories
func GetAllCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	category, err := categories.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	logging.ResponseJSON(w, category, http.StatusOK)
}

func SearchCategories(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var searchCategories = ps.ByName("search")
	categories, err := categories.SearchingCategories(ctx, searchCategories)

	if err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	logging.ResponseJSON(w, categories, http.StatusOK)
}

// Create
// PostCategories
func PostCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var category datastruct.Categories

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := categories.Insert(ctx, category); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)

}

// UpdateCategories
func UpdateCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var category datastruct.Categories

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCategories = ps.ByName("id")

	if err := categories.Update(ctx, category, idCategories); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)
}

// Delete
// DeleteCategories
func DeleteCategories(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idCategories = ps.ByName("id")

	if err := categories.Delete(ctx, idCategories); err != nil {
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