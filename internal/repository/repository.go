package repository

import (
	"context"
	"database/sql"

	"github.com/mochammadshenna/aplikasi-po/internal/entity"
)

type PurchaseOrderRepository interface {
	FindAdminByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Admin, error)
	SaveAdmin(ctx context.Context, tx *sql.Tx, admin entity.Admin) error
	ValidateAdminCredentials(ctx context.Context, tx *sql.Tx, email, password string) (entity.Admin, error)

	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.PurchaseOrder, error)
	FindById(ctx context.Context, tx *sql.Tx, poId int) (entity.PurchaseOrder, error)
	SavePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder) (entity.PurchaseOrder, error)
	UpdatePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder, poIds int64) (entity.PurchaseOrder, error)
	DeletePurchaseOrder(ctx context.Context, tx *sql.Tx, poId int64)

	FindFinishingFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.FinishingFactory, error)
	FindProductionFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.ProductionFactory, error)
}
