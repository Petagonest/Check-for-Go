package main

import (
	"Check-for-Go/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Petagonest/Check-for-Go/datastruct"
	"github.com/Petagonest/Check-for-Go/service/categories"
	"github.com/Petagonest/Check-for-Go/service/products"
	"github.com/Petagonest/Check-for-Go/service/stores"
	"github.com/julienschmidt/httprouter"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

		//storeprofile
		router := httprouter.New()
		router.GET("/stores", Auth(GetStore))
		router.POST("/stores", Auth(PostStore))
		router.PUT("/stores/update/:id", (UpdateStore))
		router.DELETE("/stores/delete/:id", (DeleteStore))
		////////////////////////////////////////////////////

		//products
		router.GET("/products", Auth(GetProducts))
		router.POST("/products", PostProducts)
		router.PUT("/products/update/:id", UpdateProducts)
		router.DELETE("/products/delete/:id", Auth(DeleteProducts))
		////////////////////////////////////////////////////

		//Categories
		router.GET("/categories", Auth(GetCategories))
		router.POST("/categories", Auth(PostCategories))
		router.PUT("/categories/update/:id", Auth(UpdateCategories))
		router.DELETE("/categories/delete/:id", Auth(DeleteCategories))
		////////////////////////////////////////////////////

		// untuk menampilkan file html di folder public
		router.NotFound = http.FileServer(http.Dir("public"))

		fmt.Println("AMAN")
		log.Fatal(http.ListenAndServe(":"+port, router))
	}
}

//auth
func Auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == "admin" && password == "admin" {
			// Delegate request to the given handle
			h(w, r, ps)
		} else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

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

	utils.ResponseJSON(w, stores, http.StatusOK)
}

// Create
// PostStore
func PostStore(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var str datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := stores.Insert(ctx, str); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateStore
func UpdateStore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var str datastruct.Stores

	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idStores = ps.ByName("id")

	if err := stores.Update(ctx, str, idStores); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
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
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

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

	utils.ResponseJSON(w, prd, http.StatusOK)
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
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := products.Insert(ctx, prd); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

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
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idProducts = ps.ByName("id")

	if err := products.Update(ctx, prd, idProducts); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
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
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////

//--------Categories----------//
// Read
// GetCategories
func GetCategories(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	ctgr, err := categories.GetAll(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(w, ctgr, http.StatusOK)
}

// Create
// PostCategories
func PostCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ctgr datastruct.Categories

	if err := json.NewDecoder(r.Body).Decode(&ctgr); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	if err := categories.Insert(ctx, ctgr); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)

}

// UpdateCategories
func UpdateCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var ctgr datastruct.Categories

	if err := json.NewDecoder(r.Body).Decode(&ctgr); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	}

	var idCategories = ps.ByName("id")

	if err := categories.Update(ctx, ctgr, idCategories); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusCreated)
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
		utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(w, res, http.StatusOK)
}

/////////////////////////////////////////////////////////////////////
