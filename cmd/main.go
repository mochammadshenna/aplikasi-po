package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/app"
	"github.com/mochammadshenna/aplikasi-po/controller"
	"github.com/mochammadshenna/aplikasi-po/repository"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

func main() {
	router := httprouter.New()
	validate := validator.New()
	db := app.NewDB()
	host := "localhost:8080"

	purchaseRepository := repository.NewPurchaseRepository()
	purchaseService := service.NewPurchaseOrderService(purchaseRepository, db, validate)
	purchaseController := controller.NewPurchaseOrderController(purchaseService)

	router.GET("/api/purchase", purchaseController.FindAllPurchaceOrder)
	router.GET("/api/purchase/:id", purchaseController.FindPurchaceOrderById)
	router.POST("/api/purchase", purchaseController.SavePurchaceOrder)
	router.PUT("/api/purchase/:id", purchaseController.UpdatePurchaceOrder)
	router.DELETE("/api/purchase/:id", purchaseController.DeletePurchaceOrder)

	router.GET("/api/produksi/:id", purchaseController.FindProductionFactory)
	router.GET("/api/finishing/:id", purchaseController.FindFinsihingFactory)

	router.PanicHandler = exceptioncode.ErrorHandler

	server := http.Server{
		Addr:    host,
		Handler: router,
	}

	fmt.Printf("Apps running on host:%s", host)

	err := server.ListenAndServe()
	helper.PanicError(err)
}
