package state

var (
	_productStatuses       productStatuses
)

func init() {
	_productStatuses = newUserStatuses()
}

type ProductStatus int64

func (s ProductStatus) Int64() int64 {
	return int64(s)
}

func (s ProductStatus) Name() string {
	switch s {
	case ProductStatuses().PENDING:
		return "PENDING"
	case ProductStatuses().DELIVERY:
		return "DELIVERY"
	case ProductStatuses().REJECTED:
		return "REJECTED"
	case ProductStatuses().SUCCESS:
		return "SUCCESS"
	default:
		return ""
	}
}

type productStatuses struct {
	PENDING  ProductStatus
	DELIVERY ProductStatus
	REJECTED ProductStatus
	SUCCESS ProductStatus
}

func newUserStatuses() productStatuses {
	return productStatuses{
		PENDING:  1,
		DELIVERY: 2,
		REJECTED: 3,
		SUCCESS: 4,
	}
}

func ProductStatuses() productStatuses {
	return _productStatuses
}
