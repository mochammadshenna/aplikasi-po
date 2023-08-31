package helper

import (
	"github.com/mochammadshenna/aplikasi-po/entity"
	"github.com/mochammadshenna/aplikasi-po/model/api"
)

func ToSavePurchaseOrderResponse(po entity.PurchaseOrder) api.SavePurchaseOrderResponse {
	return api.SavePurchaseOrderResponse{
		Success: true,
	}
}

func ToUpdatePurchaseOrderResponse(po entity.PurchaseOrder) api.UpdatePurchaseOrderResponse {
	return api.UpdatePurchaseOrderResponse{
		Success: true,
	}
}
