package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type requestBody struct {
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(tasks)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var recBody requestBody
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&recBody)
	if err != nil {
		log.Fatal(err)
	}
	newTask := Task{
		Task:   recBody.Task,
		IsDone: recBody.IsDone,
	}
	DB.Create(&newTask)
	err = json.NewEncoder(w).Encode(newTask)
	if err != nil {
		log.Println("Error encoding JSON:", err)
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func main() {
	InitDB()
	err := DB.AutoMigrate(&Task{})
	if err != nil {
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", PostHandler).Methods("POST")
	router.HandleFunc("/", GetHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
