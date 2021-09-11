package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ol1bot/base-api/auth"
	"golang.org/x/oauth2"
)

func AuthHander(r *mux.Router) {

	r.HandleFunc("",func(rw http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(rw, "Auth Route")
	})

	r.HandleFunc("/google", Google)

	r.HandleFunc("/google/callback", GoogleCallback)
}

func Google(w http.ResponseWriter, r *http.Request) {

	// auth.GoogleAuthConfig()
	conf, err := auth.GoogleAuthConfig()
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}

	url := conf.AuthCodeURL("state")

	http.Redirect(w, r, url, 302)
}

func GoogleCallback(w http.ResponseWriter, r *http.Request) {

	code := r.URL.Query().Get("code")
	conf, err := auth.GoogleAuthConfig()
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}

	// TODO 
	// store user token in database

	client := conf.Client(oauth2.NoContext, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo?alt=json")
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bt, err := ioutil.ReadAll(resp.Body)

	authedPerson := &auth.GoogleProfileBody{}

	if err = json.Unmarshal(bt, &authedPerson); err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
		return
	}

	if err = auth.ExistsInDatabase(authedPerson.Email); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to authorize"))
		return
	}

	w.Write(bt)

	// TODO
	// compare email to that in DB, grab permissions, info etc.
	// turn into JWT
	// return our token w/e information back to user
}

