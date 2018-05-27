package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	block "github.com/remoteview/service-blocks/blocks"
)

var blocks []block.Block

func main() {
	blocks = append(blocks, block.Block{ID: "1", StartTime: block.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: block.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "1"})
	blocks = append(blocks, block.Block{ID: "2", StartTime: block.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: block.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "2"})
	blocks = append(blocks, block.Block{ID: "3", StartTime: block.JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: block.JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "3"})

	router := mux.NewRouter()
	router.HandleFunc("/_health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/blocks", listBlocksHandler).Methods("GET")
	router.HandleFunc("/blocks/{id}", listBlockHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3001", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(block.HealthCheck{Status: "Ok"})
	if err != nil {

	}
}

func listBlocksHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(blocks)
}

func listBlockHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range blocks {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&block.Block{})
}
