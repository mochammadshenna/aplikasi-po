package controller

import (
	"embed"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/model/api"
	"github.com/mochammadshenna/aplikasi-po/service"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

//go:embed templates/*.gohtml
var templates embed.FS

type PurchaseOrder struct {
	PurchaseOrderService service.PoService
}

func NewPurchaseOrderController(purchaseService service.PoService) PurchaseOrderController {
	return &PurchaseOrder{
		PurchaseOrderService: purchaseService,
	}
}

func (controller *PurchaseOrder) FindAllPurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	// poResponses, err := controller.PurchaseOrderService.FindAllPurchaseOrder(request.Context())
	// helper.PanicError(err)

	var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	myTemplates.ExecuteTemplate(writer, "table.gohtml", nil)

	// webResponse := web.WebResponse{
	// 	Code:   http.StatusOK,
	// 	Status: "OK",
	// 	Data:   poResponses,
	// }

	// helper.WriteToResponseBody(writer, "table.gohtml")
}

func (controller *PurchaseOrder) FindPurchaceOrderById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := api.FindPurchaseOrderByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	poResponse := controller.PurchaseOrderService.FindPurchaseOrderById(request.Context(), poById)
	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) SavePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poSave := api.SavePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poSave)

	poResponse, err := controller.PurchaseOrderService.SavePurchaseOrder(request.Context(), poSave)
	helper.PanicError(err)

	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) UpdatePurchaceOrder(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poUpdate := api.UpdatePurchaseOrderRequest{}
	helper.ReadFromRequestBody(request, poUpdate)

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poUpdate.Id = int64(id)

	poResponse, err := controller.PurchaseOrderService.UpdatePurchaseOrder(request.Context(), poUpdate)
	helper.PanicError(err)

	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
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

	poResponse := controller.PurchaseOrderService.FindProductionFactory(request.Context(), poById)
	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PurchaseOrder) FindFinsihingFactory(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	poById := api.FindFactoryByIdRequest{}

	poId := param.ByName("id")
	id, err := strconv.Atoi(poId)
	helper.PanicError(err)

	poById.Id = int64(id)

	poResponse := controller.PurchaseOrderService.FindFinishingFactory(request.Context(), poById)
	webResponse := api.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   poResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
