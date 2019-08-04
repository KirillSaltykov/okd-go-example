package main

import (
	"encoding/json"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/mux"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:30001",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
