package main

import (
	"caturilham05/product/app"
	"caturilham05/product/helper"
	"caturilham05/product/middleware"
	"caturilham05/product/repository"
	"caturilham05/product/service"
	"net/http"

	"caturilham05/product/controller"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	helper.LoadEnv(".env")
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)
	router := app.NewRouter(productController)

	server := http.Server{
		Addr:    "127.0.0.1:3001",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
