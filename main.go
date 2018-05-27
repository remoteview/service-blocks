package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Block - working block
type Block struct {
	ID        string   `json:"id,omitempty"`
	StartTime JSONTime `json:"start_time,omitempty"`
	EndTime   JSONTime `json:"end_time,omitempty"`
	Timezone  string   `json:"timezone,omitempty"`
	UserID    string   `json:"user_id,omitempty"`
}

// Marshaler -
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// JSONTime -
type JSONTime time.Time

// MarshalJSON -
func (t JSONTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("15:04"))
	return []byte(stamp), nil
}

// HealthCheck - status
type HealthCheck struct {
	Status string `json:"status"`
}

var blocks []Block

func main() {
	blocks = append(blocks, Block{ID: "1", StartTime: JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "1"})
	blocks = append(blocks, Block{ID: "2", StartTime: JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "2"})
	blocks = append(blocks, Block{ID: "3", StartTime: JSONTime(time.Date(0, 0, 0, 9, 0, 0, 0, time.UTC)), EndTime: JSONTime(time.Date(0, 0, 0, 17, 0, 0, 0, time.UTC)), Timezone: "UTC", UserID: "3"})

	router := mux.NewRouter()
	router.HandleFunc("/_health", healthCheckHandler).Methods("GET")
	router.HandleFunc("/blocks", listBlocksHandler).Methods("GET")
	router.HandleFunc("/blocks/{id}", listBlockHandler).Methods("GET")

	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS(originsOk, headersOk, methodsOk)(router)))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(HealthCheck{Status: "Ok"})
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
	json.NewEncoder(w).Encode(&Block{})
}
