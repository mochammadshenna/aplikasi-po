package service

import (
	"context"
	"database/sql"
<<<<<<< HEAD
	"fmt"
=======
	"errors"
	"os"
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mochammadshenna/aplikasi-po/internal/entity"
	"github.com/mochammadshenna/aplikasi-po/internal/model/api"
	"github.com/mochammadshenna/aplikasi-po/internal/repository"
	"github.com/mochammadshenna/aplikasi-po/internal/util/authentication"
	"github.com/mochammadshenna/aplikasi-po/internal/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
<<<<<<< HEAD
	"github.com/mochammadshenna/aplikasi-po/internal/util/password"
=======
	"golang.org/x/exp/rand"
	"google.golang.org/api/oauth2/v1"
	"google.golang.org/api/option"
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
)

type PurchaseOrderService struct {
	PurchaseOrderRepository repository.PurchaseOrderRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewPurchaseOrderService(purchaseRepository repository.PurchaseOrderRepository, DB *sql.DB, validate *validator.Validate) Service {
	return &PurchaseOrderService{
		PurchaseOrderRepository: purchaseRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *PurchaseOrderService) Login(ctx context.Context, request api.AuthAdminRequest) (api.AuthAdminResponse, error) {
<<<<<<< HEAD
	err := service.Validate.Struct(request)
	helper.PanicOnErrorContext(ctx, err)
	var result = api.AuthAdminResponse{}

	tx, err := service.DB.Begin()
	helper.PanicOnErrorContext(ctx, err)
	defer helper.CommitOrRollback(tx)

	dataAdmin, err := service.PurchaseOrderRepository.FindAdminByEmail(ctx, tx, request.Email)
	fmt.Println(dataAdmin)
	if err != nil {
		return api.AuthAdminResponse{}, api.ErrorResponse{Code: exceptioncode.CodeInvalidCredential, Message: "Incorrect email or password"}
	}

	err = password.CheckHashPassword(request.Password, dataAdmin.Password)
	if err != nil {
		return api.AuthAdminResponse{}, api.ErrorResponse{Code: exceptioncode.CodeInvalidCredential, Message: "Incorrect email or password"}
	}

	result.Token = authentication.CreateToken(time.Minute*1440, dataAdmin.Id)
	result.Name = dataAdmin.Name
	return result, nil
=======
	// Validate request
	if err := service.Validate.Struct(request); err != nil {
		return api.AuthAdminResponse{}, err
	}

	// Start transaction
	tx, err := service.DB.Begin()
	if err != nil {
		return api.AuthAdminResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	// Validate credentials
	admin, err := service.PurchaseOrderRepository.ValidateAdminCredentials(ctx, tx, request.Email, request.Password)
	if err != nil {
		return api.AuthAdminResponse{}, err
	}

	// Generate JWT token
	token := authentication.CreateToken(time.Hour*24, admin.Id)

	return api.AuthAdminResponse{
		Token: token,
		Name:  admin.Name,
	}, nil
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
}

func (service *PurchaseOrderService) FindAllPurchaseOrder(ctx context.Context) (api.FindAllPurchaceOrderRepsonse, error) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	pos, err := service.PurchaseOrderRepository.FindAll(ctx, tx)
	if err != nil {
		logger.Errorf(ctx, "An error occurred while getting the purchase order data: %v", err)
	}

	var (
		data  = api.FindAllPurchaceOrderRepsonse{}
		items []string
	)

	for _, po := range pos {
		for _, v := range po.ProductItem {
			items = append(items, v.Name)

		}
		data.List = append(data.List, api.FindPurchaseOrderResponse{
			Id:                 po.Id,
			ProductionFactory:  po.ProductionFactoryName,
			PICName:            po.PICName,
			ProductItem:        items,
			QuantityPO:         po.QuantityPO,
			QuantityProduction: po.QuantityProduction,
			PaymentTerm:        po.PaymentTerm,
			CreatedAt:          po.CreatedAt.Format("2006-01-02"),
			ExpiredAt:          po.ExpiredAt.Format("2006-01-02"),
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
		CreatedAt:          po.CreatedAt.Format("2006-01-02"),
		ExpiredAt:          po.ExpiredAt.Format("2006-01-02"),
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

	po := entity.PurchaseOrder{
		ProductionFactoryName: request.Name,
	}

	_, err = service.PurchaseOrderRepository.SavePurchaseOrder(ctx, tx, po)
	helper.PanicOnErrorContext(ctx, err)

	return api.SavePurchaseOrderResponse{
		Success: true,
	}, nil

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

	_, err = service.PurchaseOrderRepository.UpdatePurchaseOrder(ctx, tx, po, request.Id)
	helper.PanicError(err)

	return api.UpdatePurchaseOrderResponse{
		Success: true,
	}, nil
}

func (service *PurchaseOrderService) DeletePurchaseOrder(ctx context.Context, request api.DeletePurchaseOrderRequest) (api.DeletePurchaseOrderResponse, error) {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindById(ctx, tx, int(request.Id))
	if err != nil {
		logger.Errorf(ctx, "An error occurred while delete the purchase order, error:%+v", err)
		panic(exceptioncode.NewNotFoundError(err.Error()))
	}

	service.PurchaseOrderRepository.DeletePurchaseOrder(ctx, tx, po.Id)

	return api.DeletePurchaseOrderResponse{
		Success: true,
	}, nil
}

func (service *PurchaseOrderService) FindProductionFactory(ctx context.Context, request api.FindFactoryByIdRequest) api.FindProductionFactoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	po, err := service.PurchaseOrderRepository.FindProductionFactory(ctx, tx, int(request.Id))
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
<<<<<<< HEAD
=======

func (service *PurchaseOrderService) GoogleLogin(ctx context.Context, credential string) (api.AuthAdminResponse, error) {
	err := service.Validate.Struct(credential)
	helper.PanicError(err)

	tx, err := service.DB.Begin()
	helper.PanicError(err)
	defer helper.CommitOrRollback(tx)

	// Initialize the OAuth2 service
	oauth2Service, err := oauth2.NewService(ctx, option.WithAPIKey(os.Getenv("GOOGLE_CLIENT_ID")))
	if err != nil {
		return api.AuthAdminResponse{}, errors.New("failed to initialize OAuth service")
	}

	// Verify the Google token
	tokenInfo, err := oauth2Service.Tokeninfo().IdToken(credential).Do()
	if err != nil {
		return api.AuthAdminResponse{}, errors.New("invalid token")
	}

	// Check if admin exists in database
	admin, err := service.PurchaseOrderRepository.FindAdminByEmail(ctx, tx, tokenInfo.Email)
	if err != nil {
		// Create new admin if doesn't exist
		admin = entity.Admin{
			Email:    tokenInfo.Email,
			Password: service.generateRandomPassword(),
			Role:     "admin",
			Status:   "active",
		}

		err = service.PurchaseOrderRepository.SaveAdmin(ctx, tx, admin)
		if err != nil {
			logger.Errorf(ctx, "Failed to create admin from Google login: %v", err)
			return api.AuthAdminResponse{}, errors.New("failed to create admin")
		}
	}

	// Generate JWT token
	token := authentication.CreateToken(time.Minute*1440, admin.Id)

	return api.AuthAdminResponse{
		Token: token,
		Name:  admin.Name,
	}, nil
}

// Add this helper function to generate a random password for Google-authenticated users
func (service *PurchaseOrderService) generateRandomPassword() string {
	const length = 16
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
