package routes

import (
<<<<<<< HEAD
	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/internal/app/middleware"
=======
	"net/http"

	"github.com/julienschmidt/httprouter"
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
	"github.com/mochammadshenna/aplikasi-po/internal/controller"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
)

func NewRouter(purchaseController controller.PurchaseOrderController) *httprouter.Router {
	router := httprouter.New()

<<<<<<< HEAD
	router.POST("/login", purchaseController.Login)

	router.GET("/api/purchase", middleware.JWTAuth(purchaseController.FindAllPurchaceOrder))
	router.GET("/api/purchase/:id", middleware.JWTAuth(purchaseController.FindPurchaceOrderById))
	router.POST("/api/purchase", middleware.JWTAuth(purchaseController.SavePurchaceOrder))
	router.PUT("/api/purchase/:id", middleware.JWTAuth(purchaseController.UpdatePurchaceOrder))
	router.DELETE("/api/purchase/:id", middleware.JWTAuth(purchaseController.DeletePurchaceOrder))

	router.GET("/api/production/:id", middleware.JWTAuth(purchaseController.FindProductionFactory))
	router.GET("/api/finishing/:id", middleware.JWTAuth(purchaseController.FindFinsihingFactory))

	router.PanicHandler = helper.ErrorHandler

=======
	// router.POST("/login", purchaseController.Login)

	// router.POST("/auth/google", purchaseController.GoogleLogin)

	// router.GET("/api/purchase", middleware.JWTAuth(purchaseController.FindAllPurchaceOrder))
	// router.GET("/api/purchase/:id", middleware.JWTAuth(purchaseController.FindPurchaceOrderById))
	// router.POST("/api/purchase", middleware.JWTAuth(purchaseController.SavePurchaceOrder))
	// router.PUT("/api/purchase/:id", middleware.JWTAuth(purchaseController.UpdatePurchaceOrder))
	// router.DELETE("/api/purchase/:id", middleware.JWTAuth(purchaseController.DeletePurchaceOrder))

	// router.GET("/api/production/:id", middleware.JWTAuth(purchaseController.FindProductionFactory))
	// router.GET("/api/finishing/:id", middleware.JWTAuth(purchaseController.FindFinsihingFactory))

	router.PanicHandler = helper.ErrorHandler

	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
			header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			header.Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}
		w.WriteHeader(http.StatusNoContent)
	})

>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
	return router
}
