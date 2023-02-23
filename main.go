package main

import (
	"fmt"
	"go-rest-without-framework/config"
	"go-rest-without-framework/controller"
	"go-rest-without-framework/helpers"
	"go-rest-without-framework/repository"
	"go-rest-without-framework/routers"
	"go-rest-without-framework/service"
	"net/http"
)

func main() {
	fmt.Printf("Start server")
	// database
	db := config.DatabaseConnectoon()

	// repository
	bookRepository := repository.NewBookRepositoryImpl(db)

	// service
	bookService := service.NewBookServiceImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := routers.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: routes}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
