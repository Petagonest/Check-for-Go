package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	transport_categories "github.com/Petagonest/Check-for-Go/transport/categories"
	transport_products "github.com/Petagonest/Check-for-Go/transport/products"
	transport_stores "github.com/Petagonest/Check-for-Go/transport/stores"
	"github.com/julienschmidt/httprouter"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

		//storeprofile
		router := httprouter.New()
		router.GET("/stores", Auth(transport_stores.GetAllStore))
		router.GET("/stores/:search", Auth(transport_stores.GetOneStore))
		router.POST("/stores", Auth(transport_stores.PostStore))
		router.PUT("/stores/:id", Auth(transport_stores.UpdateStore))
		router.DELETE("/stores/:id", Auth(transport_stores.DeleteStore))
		////////////////////////////////////////////////////

		//products
		router.GET("/products", Auth(transport_products.GetProducts))
		router.POST("/products", transport_products.PostProducts)
		router.PUT("/products/:id", transport_products.UpdateProducts)
		router.DELETE("/products/:id", Auth(transport_products.DeleteProducts))
		////////////////////////////////////////////////////

		//Categories
		router.GET("/categories", Auth(transport_categories.GetCategories))
		router.POST("/categories", Auth(transport_categories.PostCategories))
		router.PUT("/categories/:id", Auth(transport_categories.UpdateCategories))
		router.DELETE("/categories/:id", Auth(transport_categories.DeleteCategories))
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
