package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ol1bot/base-api/routes"
)

func main() {

	r := mux.NewRouter().StrictSlash(false)

	authRouter := r.PathPrefix("/oauth").Subrouter()
	routes.AuthHander(authRouter)
	
	
	
	
	apiRouter := r.PathPrefix("/v1/api").Subrouter()
	routes.ApiHander(apiRouter)



	r.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"Server Root")
	})
	
	http.ListenAndServe(":3000", r)

}