package repository

import (
	"context"
	"database/sql"

	"github.com/mochammadshenna/aplikasi-po/model/domain"
)

type PurchaseOrderRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PurchaseOrder, error)
	FindById(ctx context.Context, tx *sql.Tx, poId int) (domain.PurchaseOrder, error)
	SavePurchaseOrder(ctx context.Context, tx *sql.Tx, po domain.PurchaseOrder) (domain.PurchaseOrder, error)
	UpdatePurchaseOrder(ctx context.Context, tx *sql.Tx, po domain.PurchaseOrder, poIds int64) (domain.PurchaseOrder, error)
	DeletePurchaseOrder(ctx context.Context, tx *sql.Tx, poId int64)

	FindFinishingFactory(ctx context.Context, tx *sql.Tx, codeId int) (domain.FinishingFactory, error)
	FindProductionFactory(ctx context.Context, tx *sql.Tx, codeId int) (domain.ProductionFactory, error)
}
