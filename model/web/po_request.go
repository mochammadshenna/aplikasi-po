package web

type FindFactoryByIdRequest struct {
	Id int64 `validate:"required,number"`
}

type SavePurchaseOrderRequest struct {
	Name string `json:"name"`
}

type UpdatePurchaseOrderRequest struct {
	Id   int64
	Name string `validate:"required,max=200,min=1"`
}

type DeletePurchaseOrderRequest struct {
	Id int64
}

type FindPurchaseOrderByIdRequest struct {
	Id int64
}
