package web

type FindAllPurchaceOrderRepsonse struct {
	List []FindPurchaseOrderResponse
}

type FindPurchaseOrderResponse struct {
	Id                 int64    `json:"id"`
	ProductionFactory  string   `json:"production_factory"`
	PICName            string   `json:"pic_name"`
	QuantityPO         int64    `json:"quantity_po"`
	QuantityProduction int64    `json:"quantity_production"`
	ProductItem        []string `json:"product_item"`
	PaymentTerm        int64    `json:"payment_term"`
	CreatedAt          string   `json:"created_at"`
	ExpiredAt          string   `json:"expired_at"`
	UnitItem           string   `json:"unit_item"`
	Description        string   `json:"description"`
	Status             string   `json:"status"`
	FinishingFactory   string   `json:"finishing_factory"`
}

type SavePurchaseOrderResponse struct {
	Success bool `json:"success"`
}

type UpdatePurchaseOrderResponse struct {
	Success bool `json:"success"`
}

type FindProductionFactoryResponse struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type FindFinishingFactoryResponse struct {
	Id   int64  `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}
