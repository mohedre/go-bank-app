package myapp

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	//define routes
	mux := mux.NewRouter()
	//mux := http.NewServeMux()
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	//starting server
	log.Fatal(http.ListenAndServe("localhost:48234", mux))
}

