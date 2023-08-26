package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/controller"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

func NewRouter(purchaseController controller.PurchaseOrderController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/purchase", purchaseController.FindAllPurchaceOrder)
	router.GET("/api/purchase/:id", purchaseController.FindPurchaceOrderById)
	router.POST("/api/purchase", purchaseController.SavePurchaceOrder)
	router.PUT("/api/purchase/:id", purchaseController.UpdatePurchaceOrder)
	router.DELETE("/api/purchase/:id", purchaseController.DeletePurchaceOrder)

	router.GET("/api/production/:id", purchaseController.FindProductionFactory)
	router.GET("/api/finishing/:id", purchaseController.FindFinsihingFactory)

	router.PanicHandler = helper.ErrorHandler

	return router
}
