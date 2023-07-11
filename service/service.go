package service

import (
	"context"

	"github.com/mochammadshenna/aplikasi-po/model/web"
)

type PoService interface {
	FindAllPurchaseOrder(ctx context.Context) (web.FindAllPurchaceOrderRepsonse, error)
	FindPurchaseOrderById(ctx context.Context, request web.FindPurchaseOrderByIdRequest) web.FindPurchaseOrderResponse
	SavePurchaseOrder(ctx context.Context, request web.SavePurchaseOrderRequest) (res web.SavePurchaseOrderResponse, err error)
	UpdatePurchaseOrder(ctx context.Context, request web.UpdatePurchaseOrderRequest) (res web.UpdatePurchaseOrderResponse, err error)
	DeletePurchaseOrder(ctx context.Context, request web.DeletePurchaseOrderRequest)

	FindProductionFactory(ctx context.Context, request web.FindFactoryByIdRequest) web.FindProductionFactoryResponse
	FindFinishingFactory(ctx context.Context, request web.FindFactoryByIdRequest) web.FindFinishingFactoryResponse
}
