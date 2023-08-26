package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type PurchaseOrderController interface {
	FindAllPurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindPurchaceOrderById(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	SavePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	UpdatePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	DeletePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params)

	FindProductionFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
	FindFinsihingFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params)
}
