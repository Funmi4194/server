package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	h := mux.NewRouter()

	h.HandleFunc("/dashboard/{username}", func(w http.ResponseWriter, r *http.Request) {
		p := mux.Vars(r)
		username := p["username"]

		type Register struct {
			Name string `json:"name"`
		}
		a := Register{}

		json.NewDecoder(r.Body).Decode(&a)

		w.Header().Add("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{
			"name":     a.Name,
			//"method":   r.Method,
			"username": username,
		})

		//fmt.Fprint(w, "<h1>Welcome</h1>")
	}).Methods("POST")

	s := &http.Server{
		Addr:           ":8080",
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Printf("Starting server on http://localhost:%d\n", 8080)
	log.Fatal(s.ListenAndServe())
}
