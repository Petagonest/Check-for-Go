package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	transport_categories "github.com/Petagonest/Check-for-Go/transport/categories"
	transport_products "github.com/Petagonest/Check-for-Go/transport/products"
	// "github.com/Petagonest/Check-for-Go/transport/searching"
	transport_stores "github.com/Petagonest/Check-for-Go/transport/stores"
	"github.com/julienschmidt/httprouter"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

		//storeprofile
		router := httprouter.New()
		router.GET("/stores", transport_stores.GetAllStore)
		router.GET("/stores/:search", transport_stores.SearchStores)
		router.POST("/stores", Auth(transport_stores.PostStore))
		router.PUT("/stores/:id", Auth(transport_stores.UpdateStore))
		router.DELETE("/stores/:id", Auth(transport_stores.DeleteStore))
		////////////////////////////////////////////////////

		//products
		router.GET("/products", transport_products.GetAllProducts)
		router.GET("/products/:search", transport_products.SearchProducts)
		router.POST("/products", Auth(transport_products.PostProducts))
		router.PUT("/products/:id", Auth(transport_products.UpdateProducts))
		router.DELETE("/products/:id", Auth(transport_products.DeleteProducts))
		////////////////////////////////////////////////////

		//Categories
		router.GET("/categories", (transport_categories.GetAllCategories))
		router.GET("/categories/:search", transport_categories.SearchCategories)
		router.POST("/categories", Auth(transport_categories.PostCategories))
		router.PUT("/categories/:id", Auth(transport_categories.UpdateCategories))
		router.DELETE("/categories/:id", Auth(transport_categories.DeleteCategories))

		// router.GET("/searching/:search", (searching.Search))
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
