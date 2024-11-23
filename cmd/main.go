package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	config "github.com/mochammadshenna/aplikasi-po/configs"
	"github.com/mochammadshenna/aplikasi-po/internal/app/database"
	"github.com/mochammadshenna/aplikasi-po/internal/app/middleware"
	"github.com/mochammadshenna/aplikasi-po/internal/controller"
	"github.com/mochammadshenna/aplikasi-po/internal/repository"
	"github.com/mochammadshenna/aplikasi-po/internal/routes"
	"github.com/mochammadshenna/aplikasi-po/internal/service"
	"github.com/mochammadshenna/aplikasi-po/internal/state"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
)

func main() {
	config.Init(state.App.Environment)
	logger.Init()

	validate := validator.New()
	db := database.NewDb()

	purchaseRepository := repository.NewPurchaseRepository()
	purchaseService := service.NewPurchaseOrderService(purchaseRepository, db, validate)
	purchaseController := controller.NewPurchaseOrderController(purchaseService)
	router := routes.NewRouter(purchaseController)

	host := fmt.Sprintf("%s:%d", config.Get().Server.Host, config.Get().Server.Port)
	fmt.Printf("Server running on host:%d \n", config.Get().Server.Port)

	server := http.Server{
		Addr:    host,
		Handler: middleware.MultipleMiddleware(middleware.NewHttpMiddleware(router)),
	}

	err := server.ListenAndServe()
	helper.PanicError(err)
}
