package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/mochammadshenna/aplikasi-po/app"
	"github.com/mochammadshenna/aplikasi-po/controller"
	"github.com/mochammadshenna/aplikasi-po/repository"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/logger"
)

func main() {

	logger.Init()

	validate := validator.New()
	db := app.NewDB()

	purchaseRepository := repository.NewPurchaseRepository()
	purchaseService := service.NewPurchaseOrderService(purchaseRepository, db, validate)
	purchaseController := controller.NewPurchaseOrderController(purchaseService)
	router := app.NewRouter(purchaseController)

	host := "localhost:8080"
	server := http.Server{
		Addr:    host,
		Handler: router,
	}

	fmt.Printf("Server running on host:%s \n", host)

	err := server.ListenAndServe()
	helper.PanicError(err)
}
