package web

type FindPurchaseOrderByIdRequest struct {
	Id int64
}

type SavePurchaseOrderRequest struct {
	FactoryName string
}

type UpdatePurchaseOrderRequest struct {
	Id int64
}

type DeletePurchaseOrderRequest struct {
	Id int64
}
