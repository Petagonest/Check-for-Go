package transport_stores

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/logging"
	"github.com/Petagonest/Check-for-Go/service/stores"
	"github.com/julienschmidt/httprouter"
)

//------ store -----//
// Read
// GetAllStore
func GetAllStore(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	stores, err := stores.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	logging.ResponseJSON(w, stores, http.StatusOK)
}

// GetOneStore
func SearchStores(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var searchStores = ps.ByName("search")
	stores, err := stores.SearchingStores(ctx, searchStores)

	if err != nil {
		kesalahan := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		logging.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	logging.ResponseJSON(w, stores, http.StatusOK)
}

// Create
// PostStore
func PostStore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var store datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := stores.Insert(ctx, store); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)

}

// UpdateStore
func UpdateStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var store datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&store); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idStores = ps.ByName("id")

	if err := stores.Update(ctx, store, idStores); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsukses := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, responsukses, http.StatusCreated)
}

// Delete
// DeleteStore
func DeleteStore(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idStores = ps.ByName("id")

	if err := stores.Delete(ctx, idStores); err != nil {
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
