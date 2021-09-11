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

	// r.HandleFunc("/portal/", func(rw http.ResponseWriter, r *http.Request) {})
	// r.HandleFunc("/students/", func(rw http.ResponseWriter, r *http.Request) {})
	// r.HandleFunc("/student_years/:year", func(rw http.ResponseWriter, r *http.Request) {})
	// r.HandleFunc("/staff", func(rw http.ResponseWriter, r *http.Request) {})
	// r.HandleFunc("/payees", func(rw http.ResponseWriter, r *http.Request) {})


}