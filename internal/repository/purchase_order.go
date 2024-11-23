package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mochammadshenna/aplikasi-po/internal/entity"
	"github.com/mochammadshenna/aplikasi-po/internal/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
)

type PurchaseOrderRepository interface {
	FindAdminByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Admin, error)

	FindAll(ctx context.Context, tx *sql.Tx) ([]entity.PurchaseOrder, error)
	FindById(ctx context.Context, tx *sql.Tx, poId int) (entity.PurchaseOrder, error)
	SavePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder) (entity.PurchaseOrder, error)
	UpdatePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder, poIds int64) (entity.PurchaseOrder, error)
	DeletePurchaseOrder(ctx context.Context, tx *sql.Tx, poId int64)

	FindFinishingFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.FinishingFactory, error)
	FindProductionFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.ProductionFactory, error)
}

type repository struct {
}

func NewPurchaseRepository() PurchaseOrderRepository {
	return &repository{}
}

func (repo *repository) FindAdminByEmail(ctx context.Context, tx *sql.Tx, email string) (entity.Admin, error) {
	query := `SELECT id, name, created_at, updated_at, email, password FROM admins WHERE email = $1`
	var result entity.Admin

	err := tx.QueryRowContext(ctx, query, email).Scan(
		&result.Id,
		&result.Name,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.Email,
		&result.Password,
	)

	if err == sql.ErrNoRows {
		return result, exceptioncode.ErrEmptyResult
	}
	helper.PanicOnErrorContext(ctx, err)

	return result, nil
}

func (repository *repository) FindAll(ctx context.Context, tx *sql.Tx) ([]entity.PurchaseOrder, error) {
	query := `SELECT
				po.id,
				pf.name,
				po.pic_name,
				po.quantity_po,
				po.quantity_production,
				po.product_item,
				po.payment_term,
				po.created_at,
				po.expired_at,
				po.unit_item,
				po.description,
				po.status,
				po.status_history,
				ff.name
			FROM purchase_orders po
			LEFT JOIN production_factories pf ON pf.id = po.production_factory
			LEFT JOIN finishing_factories ff ON ff.id = po.finishing_factory`

	var result []entity.PurchaseOrder
	rows, err := tx.QueryContext(ctx, query)
	helper.TranslatePostgreError(ctx, err)
	defer func() {
		err = rows.Close()
		helper.PanicOnErrorContext(ctx, err)
	}()
	for rows.Next() {
		var po entity.PurchaseOrder
		err := rows.Scan(
			&po.Id,
			&po.ProductionFactoryName,
			&po.PICName,
			&po.QuantityPO,
			&po.QuantityProduction,
			&po.ProductItem,
			&po.PaymentTerm,
			&po.CreatedAt,
			&po.ExpiredAt,
			&po.UnitItem,
			&po.Description,
			&po.Status,
			&po.StatusHistory,
			&po.FinishingFactoryName,
		)
		helper.TranslatePostgreError(ctx, err)
		result = append(result, po)
	}

	if len(result) == 0 {
		return result, errors.New("empty result")
	}

	return result, nil
}

func (repository *repository) FindById(ctx context.Context, tx *sql.Tx, poId int) (entity.PurchaseOrder, error) {
	query := `SELECT
				po.id,
				pf.name,
				po.pic_name,
				po.quantity_po,
				po.quantity_production,
				po.product_item,
				po.payment_term,
				po.created_at,
				po.expired_at,
				po.unit_item,
				po.description,
				po.status,
				po.status_history,
				ff.name
			FROM purchase_orders po
			LEFT JOIN production_factories pf ON pf.id = po.production_factory
			LEFT JOIN finishing_factories ff ON ff.id = po.finishing_factory
			WHERE po.id = $1`

	var po entity.PurchaseOrder

	rows, err := tx.QueryContext(ctx, query, poId)
	helper.TranslatePostgreError(ctx, err)
	defer func() {
		err = rows.Close()
		helper.PanicOnErrorContext(ctx, err)
	}()

	if rows.Next() {
		err := rows.Scan(
			&po.Id,
			&po.ProductionFactoryName,
			&po.PICName,
			&po.QuantityPO,
			&po.QuantityProduction,
			&po.ProductItem,
			&po.PaymentTerm,
			&po.CreatedAt,
			&po.ExpiredAt,
			&po.UnitItem,
			&po.Description,
			&po.Status,
			&po.StatusHistory,
			&po.FinishingFactoryName,
		)
		helper.TranslatePostgreError(ctx, err)
		return po, nil
	} else {
		return po, errors.New("PO is not found")
	}
}

func (repository *repository) SavePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder) (entity.PurchaseOrder, error) {
	query := `INSERT INTO purchase_orders(
					production_factory,
					pic_name,
					quantity_po,
					quantity_production,
					product_item,
					payment_term,
					created_at,
					expired_at,
					unit_item,
					description,
					status,
					status_history,
					finishing_factory)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
		RETURNING id`

	err := tx.QueryRowContext(ctx, query,
		&po.ProductionFactory,
		&po.PICName,
		&po.QuantityPO,
		&po.QuantityProduction,
		&po.ProductItem,
		&po.PaymentTerm,
		&po.CreatedAt,
		&po.ExpiredAt,
		&po.UnitItem,
		&po.Description,
		&po.Status,
		&po.StatusHistory,
		&po.FinishingFactory).Scan(&po.Id)
	helper.TranslatePostgreError(ctx, err)

	return po, nil
}

func (repository *repository) UpdatePurchaseOrder(ctx context.Context, tx *sql.Tx, po entity.PurchaseOrder, poIds int64) (entity.PurchaseOrder, error) {
	query := `UPDATE purchase_orders
		SET
			production_factory=$1,
			pic_name=$2,
			quantity_po=$3,
			quantity_production=$4,
			product_item=$5,
			payment_term=$6,
			created_at=$7,
			expired_at=$8,
			unit_item=$9,
			description=$10,
			status=$11,
			status_history=$12,
			finishing_factory=$13`

	res, err := tx.ExecContext(ctx, query,
		&po.ProductionFactory,
		&po.PICName,
		&po.QuantityPO,
		&po.QuantityProduction,
		&po.ProductItem,
		&po.PaymentTerm,
		&po.CreatedAt,
		&po.ExpiredAt,
		&po.UnitItem,
		&po.Description,
		&po.Status,
		&po.StatusHistory,
		&po.FinishingFactory)
	helper.TranslatePostgreError(ctx, err)
	r, err := res.RowsAffected()
	helper.TranslatePostgreError(ctx, err)
	if r == 0 {
		return po, exceptioncode.ErrEmptyResult
	}

	return po, nil
}

func (repository *repository) DeletePurchaseOrder(ctx context.Context, tx *sql.Tx, poId int64) {
	sql := "DELETE FROM purchase_orders WHERE id = $1"

	_, err := tx.ExecContext(ctx, sql, poId)
	helper.TranslatePostgreError(ctx, err)
}

func (repository *repository) FindFinishingFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.FinishingFactory, error) {
	query := `SELECT
				id,
				code,
				name
			FROM finishing_factories
			WHERE id = $1`

	var pc entity.FinishingFactory

	rows, err := tx.QueryContext(ctx, query, codeId)
	if err != nil {
		helper.TranslatePostgreError(ctx, err)
	}
	defer func() {
		err = rows.Close()
		helper.PanicOnErrorContext(ctx, err)
	}()

	if rows.Next() {
		err := rows.Scan(
			&pc.Id,
			&pc.Code,
			&pc.Name,
		)
		helper.TranslatePostgreError(ctx, err)
		return pc, nil
	} else {
		return pc, errors.New("PO is not found")
	}
}

func (repository *repository) FindProductionFactory(ctx context.Context, tx *sql.Tx, codeId int) (entity.ProductionFactory, error) {
	query := `SELECT
				id,
				name
			FROM production_factories
			WHERE id = $1`

	var pc entity.ProductionFactory

	rows, err := tx.QueryContext(ctx, query, codeId)
	if err != nil {
		helper.TranslatePostgreError(ctx, err)
	}
	defer func() {
		err = rows.Close()
		helper.PanicOnErrorContext(ctx, err)
	}()

	if rows.Next() {
		err := rows.Scan(
			&pc.Id,
			&pc.Name,
		)
		helper.TranslatePostgreError(ctx, err)
		return pc, nil
	} else {
		return pc, errors.New("PO is not found")
	}
}
