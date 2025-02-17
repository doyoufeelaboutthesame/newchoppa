package main

import (
	"TheRealOne/internal/database"
	"TheRealOne/internal/handlers"
	"TheRealOne/internal/taskService"
	"TheRealOne/internal/web/tasks"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
