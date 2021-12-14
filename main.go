package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/Petagonest/Check-for-Go/transport/stores"
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
