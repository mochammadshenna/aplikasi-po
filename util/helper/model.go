package helper

import (
	"github.com/mochammadshenna/aplikasi-po/model/domain"
	"github.com/mochammadshenna/aplikasi-po/model/web"
)

func ToSavePurchaseOrderResponse(po domain.PurchaseOrder) web.SavePurchaseOrderResponse {
	return web.SavePurchaseOrderResponse{
		Success: true,
	}
}

func ToUpdatePurchaseOrderResponse(po domain.PurchaseOrder) web.UpdatePurchaseOrderResponse {
	return web.UpdatePurchaseOrderResponse{
		Success: true,
	}
}
