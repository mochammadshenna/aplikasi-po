package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/mochammadshenna/aplikasi-po/app"
	config "github.com/mochammadshenna/aplikasi-po/configs"
	"github.com/mochammadshenna/aplikasi-po/controller"
	"github.com/mochammadshenna/aplikasi-po/repository"
	"github.com/mochammadshenna/aplikasi-po/routes"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/state"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/logger"
)

func main() {
	config.Init(state.App.Environment)
	logger.Init()

	validate := validator.New()
	db := app.NewDb()

	purchaseRepository := repository.NewPurchaseRepository()
	purchaseService := service.NewPurchaseOrderService(purchaseRepository, db, validate)
	purchaseController := controller.NewPurchaseOrderController(purchaseService)
	router := routes.NewRouter(purchaseController)

	host := fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port)
	fmt.Printf("Server running on host:%d \n", config.Get().Server.Port)

	server := http.Server{
		Addr:    host,
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
