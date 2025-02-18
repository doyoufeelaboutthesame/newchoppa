package main

import (
	"TheRealOne/internal/database"
	"TheRealOne/internal/handlers"
	"TheRealOne/internal/taskService"
	"TheRealOne/internal/userService"
	"TheRealOne/internal/web/tasks"
	"TheRealOne/internal/web/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	database.InitDB()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//===================================================

	err := database.DB.AutoMigrate(&taskService.Task{})
	if err != nil {
		log.Fatal(err)
	}
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewTaskService(repo)
	handler := handlers.NewHandler(service)

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	//===================================================
	err = database.DB.AutoMigrate(&userService.User{})
	if err != nil {
		log.Fatal(err)
	}
	uRepo := userService.NewUserRepository(database.DB)
	uService := userService.NewUserService(uRepo)
	uHandler := handlers.NewUserHandler(uService)

	uStrictHandler := users.NewStrictHandler(uHandler, nil)
	users.RegisterHandlers(e, uStrictHandler)
	//===================================================
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
