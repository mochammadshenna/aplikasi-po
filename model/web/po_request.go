package web

type FindFactoryByIdRequest struct {
	Id int64
}

type SavePurchaseOrderRequest struct {
	ProductionFactoryName string
}

type UpdatePurchaseOrderRequest struct {
	Id int64
}

type DeletePurchaseOrderRequest struct {
	Id int64
}

type FindPurchaseOrderByIdRequest struct {
	Id int64
}
