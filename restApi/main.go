package main

import (
	"net/http"
	"restApi/app"
	"restApi/controller"
	"restApi/helper"
	"restApi/repository"
	"restApi/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {

	db := app.NewDb()
	validate := validator.New()
	categoryRepo := repository.NewCategoryRepo()
	categoryServ := service.NewCategoryService(categoryRepo, db, validate)
	categoryController := controller.NewCategoryController(categoryServ)

	router := httprouter.New()

	router.GET("/ping", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("pong"))
	})

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
