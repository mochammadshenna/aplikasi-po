package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type PurchaseOrder struct {
	Id                    int64
	ProductionFactory     int64
	PICName               string
	QuantityPO            int64
	QuantityProduction    int64
	ProductItem           ProductItems
	PaymentTerm           int64
	CreatedAt             time.Time
	ExpiredAt             time.Time
	UnitItem              string
	Description           string
	Status                string
	StatusHistory         StatHistories
	FinishingFactory      int64
	ProductionFactoryName string
	FinishingFactoryName  string
}

type ProductItems []ProdItem

type ProdItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (p ProductItems) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (s *ProductItems) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}

type StatHistories []StatHistory

type StatHistory struct {
	Status    int64     `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s StatHistories) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *StatHistories) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &s)
}

type ProductionFactory struct {
	Id   int64
	Name string
}

type FinishingFactory struct {
	Id   int64
	Code string
	Name string
}

type Admin struct {
	Id        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
}
