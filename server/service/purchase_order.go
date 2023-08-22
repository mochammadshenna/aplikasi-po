package service

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mochammadshenna/aplikasi-po/model/api"
	"github.com/mochammadshenna/aplikasi-po/model/domain"
	"github.com/mochammadshenna/aplikasi-po/repository"
	"github.com/mochammadshenna/aplikasi-po/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
	"github.com/mochammadshenna/aplikasi-po/util/logger"
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

func (service *PurchaseOrderService) FindAllPurchaseOrder(ctx context.Context) (api.FindAllPurchaceOrderRepsonse, error) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	pos, err := service.PurchaseOrderRepository.FindAll(ctx, tx)
	if err != nil {
		logger.Errorf(ctx, "An error occurred while getting the purchase order data: %v", err)
	}

	var data api.FindAllPurchaceOrderRepsonse

	for _, po := range pos {
		data.List = append(data.List, api.FindPurchaseOrderResponse{
			Id:                 po.Id,
			ProductionFactory:  po.ProductionFactoryName,
			PICName:            po.PICName,
			QuantityPO:         po.QuantityPO,
			QuantityProduction: po.QuantityProduction,
			PaymentTerm:        po.PaymentTerm,
			CreatedAt:          po.CreatedAt.Format(time.RFC3339),
			ExpiredAt:          po.ExpiredAt.Format(time.RFC3339),
			UnitItem:           po.UnitItem,
			Description:        po.Description,
			Status:             po.Status,
			FinishingFactory:   po.FinishingFactoryName,
		})
	}

	return data, nil
}

func (service *PurchaseOrderService) FindPurchaseOrderById(ctx context.Context, request api.FindPurchaseOrderByIdRequest) api.FindPurchaseOrderResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while getting the purchase order by ID: %d with error: %+v", request.Id, err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	res := api.FindPurchaseOrderResponse{
		Id:                 po.Id,
		ProductionFactory:  po.ProductionFactoryName,
		PICName:            po.PICName,
		QuantityPO:         po.QuantityPO,
		QuantityProduction: po.QuantityProduction,
		PaymentTerm:        po.PaymentTerm,
		CreatedAt:          po.CreatedAt.Format(time.RFC3339),
		ExpiredAt:          po.ExpiredAt.Format(time.RFC3339),
		UnitItem:           po.UnitItem,
		Description:        po.Description,
		Status:             po.Status,
		FinishingFactory:   po.FinishingFactoryName,
	}

	return res
}

func (service *PurchaseOrderService) SavePurchaseOrder(ctx context.Context, request api.SavePurchaseOrderRequest) (api.SavePurchaseOrderResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po := domain.PurchaseOrder{
		ProductionFactoryName: request.Name,
	}

	p, err := service.PurchaseOrderRepository.SavePurchaseOrder(ctx, tx, po)
	helper.PanicOnErrorContext(ctx, err)

	return helper.ToSavePurchaseOrderResponse(p), nil

}

func (service *PurchaseOrderService) UpdatePurchaseOrder(ctx context.Context, request api.UpdatePurchaseOrderRequest) (res api.UpdatePurchaseOrderResponse, err error) {
	errs := service.Validate.Struct(request)
	helper.PanicError(errs)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while updating the purchase order, error:%+v", err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	p, err := service.PurchaseOrderRepository.UpdatePurchaseOrder(ctx, tx, po, request.Id)
	helper.PanicError(err)

	return helper.ToUpdatePurchaseOrderResponse(p), nil
}

func (service *PurchaseOrderService) DeletePurchaseOrder(ctx context.Context, request api.DeletePurchaseOrderRequest) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while delete the purchase order, error:%+v", err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	service.PurchaseOrderRepository.DeletePurchaseOrder(ctx, tx, po.Id)
}

func (service *PurchaseOrderService) FindProductionFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindProductionFactoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindFinishingFactory(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while getting the production factory, error:%+v", err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	res := api.FindProductionFactoryResponse{
		Id:   po.Id,
		Name: po.Name,
	}

	logger.Info(ctx, "Successfully get production factory")
	return res
}

func (service *PurchaseOrderService) FindFinishingFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindFinishingFactoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindFinishingFactory(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while getting the finishing factory, error:%+v", err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	res := api.FindFinishingFactoryResponse{
		Id:   po.Id,
		Code: po.Code,
		Name: po.Name,
	}

	logger.Info(ctx, "Successfully get finishing factory")
	return res
}
