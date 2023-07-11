package domain

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type PurchaseOrder struct {
	Id                 int64
	FactoryName        string
	PICName            string
	QuantityPO         int64
	QuantityProduction int64
	Item               string
	PaymentTerm        int64
	CreatedAt          time.Time
	ExpiredAt          time.Time
	UnitItem           string
	Description        string
	Note               string
	Status             string
	StatusHistory      StatHistories
	PoCodeId           int64
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

type POCode struct {
	Id     int64
	CodeId int64
	Name   string
}
