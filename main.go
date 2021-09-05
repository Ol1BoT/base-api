package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	apiRoute := r.PathPrefix("/v1/api").Subrouter()


	apiRoute.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"Testing API Route")
	})

	r.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"Server Root")
	})
	
	http.ListenAndServe(":3000", r)

}