package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/model/web"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

type PurchaseOrder struct {
	PurchaseOrderService service.PoService
}

func NewPurchaseOrderController(purchaseService service.PoService) PurchaseOrderController {
	return &PurchaseOrder{
		PurchaseOrderService: purchaseService,
	}
}

func (controller *PurchaseOrder) FindAllPurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poResponses, err := controller.PurchaseOrderService.FindAllPurchaseOrder(request.Context())
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) FindPurchaceOrderById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := web.FindPurchaseOrderByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	poResponse := controller.PurchaseOrderService.FindPurchaseOrderById(request.Context(), poById)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) SavePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poSave := web.SavePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poSave)

	poResponse, err := controller.PurchaseOrderService.SavePurchaseOrder(request.Context(), poSave)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) UpdatePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poUpdate := web.UpdatePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poUpdate)

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poUpdate.Id = int64(id)

	poResponse, err := controller.PurchaseOrderService.UpdatePurchaseOrder(request.Context(), poUpdate)
	helper.PanicError(err)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) DeletePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poDelete := web.DeletePurchaseOrderRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poDelete.Id = int64(id)

	controller.PurchaseOrderService.DeletePurchaseOrder(request.Context(), poDelete)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) FindProductionFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := web.FindFactoryByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	poResponse := controller.PurchaseOrderService.FindProductionFactory(request.Context(), poById)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) FindFinsihingFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := web.FindFactoryByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	poResponse := controller.PurchaseOrderService.FindFinishingFactory(request.Context(), poById)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
