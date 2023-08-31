package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/model/api"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/httphelper"
)

type PurchaseOrder struct {
	PurchaseOrderService service.PoService
}

func NewPurchaseOrderController(purchaseService service.PoService) PurchaseOrderController {
	return &PurchaseOrder{
		PurchaseOrderService: purchaseService,
	}
}

func (controller *PurchaseOrder) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	requestData := api.AuthAdminRequest{}
	httphelper.Read(request, &requestData)

	response, errorResponse := controller.PurchaseOrderService.Login(request.Context(), requestData)
	if errorResponse != nil {
		httphelper.WriteError(request.Context(), writer, errorResponse)
		return
	}
	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) FindAllPurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	response, err := controller.PurchaseOrderService.FindAllPurchaseOrder(request.Context())
	if err != nil {
		httphelper.WriteError(request.Context(), writer, err)
		return
	}
	
	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) FindPurchaceOrderById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := api.FindPurchaseOrderByIdRequest{}
	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	response := controller.PurchaseOrderService.FindPurchaseOrderById(request.Context(), poById)
	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) SavePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poSave := api.SavePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poSave)

	response, err := controller.PurchaseOrderService.SavePurchaseOrder(request.Context(), poSave)
	if err != nil {
		httphelper.WriteError(request.Context(), writer, err)
		return
	}

	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) UpdatePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poUpdate := api.UpdatePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poUpdate)

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poUpdate.Id = int64(id)

	response, err := controller.PurchaseOrderService.UpdatePurchaseOrder(request.Context(), poUpdate)
	if err != nil {
		httphelper.WriteError(request.Context(), writer, err)
		return
	}

	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) DeletePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poDelete := api.DeletePurchaseOrderRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poDelete.Id = int64(id)

	controller.PurchaseOrderService.DeletePurchaseOrder(request.Context(), poDelete)

	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) FindProductionFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := api.FindFactoryByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	response := controller.PurchaseOrderService.FindProductionFactory(request.Context(), poById)

	httphelper.Write(request.Context(), writer, response)
}

func (controller *PurchaseOrder) FindFinsihingFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := api.FindFactoryByIdRequest{}
	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	response := controller.PurchaseOrderService.FindFinishingFactory(request.Context(), poById)

	httphelper.Write(request.Context(), writer, response)
}
