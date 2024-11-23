package service

import (
	"context"

	"github.com/mochammadshenna/aplikasi-po/internal/model/api"
)

type Service interface {
	Login(ctx context.Context, request api.AuthAdminRequest) (api.AuthAdminResponse, error)
	GoogleLogin(ctx context.Context, credential string) (api.AuthAdminResponse, error)

	FindAllPurchaseOrder(ctx context.Context) (api.FindAllPurchaceOrderRepsonse, error)
	FindPurchaseOrderById(ctx context.Context, request api.FindPurchaseOrderByIdRequest) api.FindPurchaseOrderResponse
	SavePurchaseOrder(ctx context.Context, request api.SavePurchaseOrderRequest) (api.SavePurchaseOrderResponse, error)
	UpdatePurchaseOrder(ctx context.Context, request api.UpdatePurchaseOrderRequest) (res api.UpdatePurchaseOrderResponse, err error)
	DeletePurchaseOrder(ctx context.Context, request api.DeletePurchaseOrderRequest) (api.DeletePurchaseOrderResponse, error)

	FindProductionFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindProductionFactoryResponse
	FindFinishingFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindFinishingFactoryResponse
}
