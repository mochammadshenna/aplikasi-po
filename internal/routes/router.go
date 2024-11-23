package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/internal/app/middleware"
	"github.com/mochammadshenna/aplikasi-po/internal/controller"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
)

func NewRouter(purchaseController controller.PurchaseOrderController) *httprouter.Router {
	router := httprouter.New()

	router.POST("/login", purchaseController.Login)

	router.GET("/api/purchase", middleware.JWTAuth(purchaseController.FindAllPurchaceOrder))
	router.GET("/api/purchase/:id", middleware.JWTAuth(purchaseController.FindPurchaceOrderById))
	router.POST("/api/purchase", middleware.JWTAuth(purchaseController.SavePurchaceOrder))
	router.PUT("/api/purchase/:id", middleware.JWTAuth(purchaseController.UpdatePurchaceOrder))
	router.DELETE("/api/purchase/:id", middleware.JWTAuth(purchaseController.DeletePurchaceOrder))

	router.GET("/api/production/:id", middleware.JWTAuth(purchaseController.FindProductionFactory))
	router.GET("/api/finishing/:id", middleware.JWTAuth(purchaseController.FindFinsihingFactory))

	router.PanicHandler = helper.ErrorHandler

	return router
}
