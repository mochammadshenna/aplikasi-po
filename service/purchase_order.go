package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mochammadshenna/aplikasi-po/model/domain"
	"github.com/mochammadshenna/aplikasi-po/model/web"
	"github.com/mochammadshenna/aplikasi-po/repository"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

type PurchaseOrderService struct {
	PurchaseOrderRepository repository.PurchaseOrderRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewPurchaseOrderService(purchaseRepository repository.PurchaseOrderRepository, DB *sql.DB, validate *validator.Validate) PoService {
	return &PurchaseOrderService{
		PurchaseOrderRepository: purchaseRepository,
		DB:                      DB,
		Validate:                validate,
	}

}

func (service *PurchaseOrderService) FindAllPurchaseOrder(ctx context.Context) (web.FindAllPurchaceOrderRepsonse, error) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	pos, err := service.PurchaseOrderRepository.FindAll(ctx, tx)
	helper.PanicError(err)

	var data web.FindAllPurchaceOrderRepsonse

	for _, po := range pos {
		data.List = append(data.List, web.FindPurchaseOrderResponse{
			Id:                 po.Id,
			ProductionFactory:  po.ProductionFactoryName,
			PICName:            po.PICName,
			QuantityPO:         po.QuantityPO,
			QuantityProduction: po.QuantityProduction,
			// ProductItem:        po.ProductItem,
			PaymentTerm:      po.PaymentTerm,
			CreatedAt:        po.CreatedAt.Format(time.RFC3339),
			ExpiredAt:        po.ExpiredAt.Format(time.RFC3339),
			UnitItem:         po.UnitItem,
			Description:      po.Description,
			Status:           po.Status,
			FinishingFactory: po.FinishingFactoryName,
		})
		// res = append(res, data)
	}

	return data, nil
}

func (service *PurchaseOrderService) FindPurchaseOrderById(ctx context.Context, request web.FindPurchaseOrderByIdRequest) web.FindPurchaseOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	helper.PanicError(err)

	res := web.FindPurchaseOrderResponse{
		Id:                 po.Id,
		ProductionFactory:  po.ProductionFactoryName,
		PICName:            po.PICName,
		QuantityPO:         po.QuantityPO,
		QuantityProduction: po.QuantityProduction,
		// ProductItem:        po.ProductItem,
		PaymentTerm:      po.PaymentTerm,
		CreatedAt:        po.CreatedAt.Format(time.RFC3339),
		ExpiredAt:        po.ExpiredAt.Format(time.RFC3339),
		UnitItem:         po.UnitItem,
		Description:      po.Description,
		Status:           po.Status,
		FinishingFactory: po.FinishingFactoryName,
	}

	return res
}

func (service *PurchaseOrderService) SavePurchaseOrder(ctx context.Context, request web.SavePurchaseOrderRequest) (res web.SavePurchaseOrderResponse, err error) {

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po := domain.PurchaseOrder{
		ProductionFactoryName: request.ProductionFactoryName,
	}

	p, _ := service.PurchaseOrderRepository.SavePurchaseOrder(ctx, tx, po)

	return helper.ToSavePurchaseOrderResponse(p), nil

}

func (service *PurchaseOrderService) UpdatePurchaseOrder(ctx context.Context, request web.UpdatePurchaseOrderRequest) (res web.UpdatePurchaseOrderResponse, err error) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	helper.PanicError(err)

	p, err := service.PurchaseOrderRepository.UpdatePurchaseOrder(ctx, tx, po, request.Id)
	helper.PanicError(err)

	return helper.ToUpdatePurchaseOrderResponse(p), nil
}

func (service *PurchaseOrderService) DeletePurchaseOrder(ctx context.Context, request web.DeletePurchaseOrderRequest) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	helper.PanicError(err)

	service.PurchaseOrderRepository.DeletePurchaseOrder(ctx, tx, po.Id)
}

func (service *PurchaseOrderService) FindProductionFactory(ctx context.Context, request web.FindFactoryByIdRequest) web.FindProductionFactoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindFinishingFactory(ctx, tx, int(request.Id))
	helper.PanicError(err)

	res := web.FindProductionFactoryResponse{
		Id:   po.Id,
		Name: po.Name,
	}

	return res
}

func (service *PurchaseOrderService) FindFinishingFactory(ctx context.Context, request web.FindFactoryByIdRequest) web.FindFinishingFactoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindFinishingFactory(ctx, tx, int(request.Id))
	helper.PanicError(err)

	res := web.FindFinishingFactoryResponse{
		Id:   po.Id,
		Code: po.Code,
		Name: po.Name,
	}

	return res
}
