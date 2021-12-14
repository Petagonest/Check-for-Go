package trans_stores

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
// GETstore
func GetStore(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	stores, err := stores.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	logging.ResponseJSON(w, stores, http.StatusOK)
}

// Create
// PostStore
func PostStore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var str datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := stores.Insert(ctx, str); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateStore
func UpdateStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var str datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		logging.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idStores = ps.ByName("id")

	if err := stores.Update(ctx, str, idStores); err != nil {
		logging.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusCreated)
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

	res := map[string]string{
		"status": "Succesfully",
	}

	logging.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////
