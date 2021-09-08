package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ApiHander(r *mux.Router) {

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(rw,"API Route")
	})

}