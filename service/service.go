package service

import (
	"context"

	"github.com/mochammadshenna/aplikasi-po/model/api"
)

type PoService interface {
	FindAllPurchaseOrder(ctx context.Context) (api.FindAllPurchaceOrderRepsonse, error)
	FindPurchaseOrderById(ctx context.Context, request api.FindPurchaseOrderByIdRequest) api.FindPurchaseOrderResponse
	SavePurchaseOrder(ctx context.Context, request api.SavePurchaseOrderRequest) (api.SavePurchaseOrderResponse, error)
	UpdatePurchaseOrder(ctx context.Context, request api.UpdatePurchaseOrderRequest) (res api.UpdatePurchaseOrderResponse, err error)
	DeletePurchaseOrder(ctx context.Context, request api.DeletePurchaseOrderRequest)

	FindProductionFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindProductionFactoryResponse
	FindFinishingFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindFinishingFactoryResponse
}
