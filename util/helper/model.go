package helper

import (
	"github.com/mochammadshenna/aplikasi-po/model/api"
	"github.com/mochammadshenna/aplikasi-po/model/domain"
)

func ToSavePurchaseOrderResponse(po domain.PurchaseOrder) api.SavePurchaseOrderResponse {
	return api.SavePurchaseOrderResponse{
		Success: true,
	}
}

func ToUpdatePurchaseOrderResponse(po domain.PurchaseOrder) api.UpdatePurchaseOrderResponse {
	return api.UpdatePurchaseOrderResponse{
		Success: true,
	}
}
