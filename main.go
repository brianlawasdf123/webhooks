package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.
		Methods("POST").
		Path("/payload").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			var t map[string]interface{}
			err := decoder.Decode(&t)
			if err != nil {
				panic(err)
			}
			defer r.Body.Close()

			log.Printf("%#v", t)

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
		})

	http.ListenAndServe(":8080", router)
}
