package main

import (
	"TheRealOne/internal/database"
	"TheRealOne/internal/handlers"
	"TheRealOne/internal/taskService"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	database.InitDB()
	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)

	handler := handlers.NewHandler(service)
	router := mux.NewRouter()
	router.HandleFunc("/", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/{id}", handler.PatchTaskHandler).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", router))
}
