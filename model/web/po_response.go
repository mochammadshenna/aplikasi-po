package web

type FindAllPurchaceOrderRepsonse struct {
	List []FindPurchaseOrderResponse
}

type FindPurchaseOrderResponse struct {
	Id                 int64  `json:"id"`
	FactoryName        string `json:"factory_name"`
	PICName            string `json:"pic_name"`
	QuantityPO         int64  `json:"quantity_po"`
	QuantityProduction int64  `json:"quantity_production"`
	Item               string `json:"item"`
	PaymentTerm        int64  `json:"payment_term"`
	CreatedAt          string `json:"created_at"`
	ExpiredAt          string `json:"expired_at"`
	UnitItem           string `json:"unit_item"`
	Description        string `json:"description"`
	Note               string `json:"note"`
	Status             string `json:"status"`
	PoCodeId           int64  `json:"po_code_id"`
}

type SavePurchaseOrderResponse struct {
	Success bool `json:"success"`
}

type UpdatePurchaseOrderResponse struct {
	Success bool `json:"success"`
}

type FindPOCodeResponse struct {
	Id     int64  `json:"id"`
	CodeId int64  `json:"code_id"`
	Name   string `json:"name"`
}
