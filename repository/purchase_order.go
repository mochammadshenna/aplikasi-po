package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mochammadshenna/aplikasi-po/model/domain"
	"github.com/mochammadshenna/aplikasi-po/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/util/helper"
)

type PurchaseOrder struct {
}

func NewPurchaseRepository() PurchaseOrderRepository {
	return &PurchaseOrder{}
}

func (repository *PurchaseOrder) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.PurchaseOrder, error) {
	query := `SELECT
				id,
				factory_name,
				pic_name,
				quantity_po,
				quantity_production,
				item,
				payment_term,
				created_at,
				expired_at,
				unit_item,
				description,
				note,
				status,
				status_history,
				po_code_id
			FROM purchase_orders po`

	var result []domain.PurchaseOrder
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicOnErrorContext(ctx, err)
	defer func() {
		err = rows.Close()
		helper.PanicOnErrorContext(ctx, err)
	}()
	for rows.Next() {
		var po domain.PurchaseOrder
		err := rows.Scan(
			&po.Id,
			&po.FactoryName,
			&po.PICName,
			&po.QuantityPO,
			&po.QuantityProduction,
			&po.Item,
			&po.PaymentTerm,
			&po.CreatedAt,
			&po.ExpiredAt,
			&po.UnitItem,
			&po.Description,
			&po.Note,
			&po.Status,
			&po.StatusHistory,
			&po.PoCodeId)
		helper.PanicOnErrorContext(ctx, err)
		result = append(result, po)
	}

	if len(result) == 0 {
		return result, errors.New("empty result")
	}

	return result, nil
}

func (repository *PurchaseOrder) FindById(ctx context.Context, tx *sql.Tx, poId int) (domain.PurchaseOrder, error) {
	query := `SELECT
				id,
				factory_name,
				pic_name,
				quantity_po,
				quantity_production,
				item,
				payment_term,
				created_at,
				expired_at,
				unit_item,
				description,
				note,
				status,
				status_history,
				po_code_id
			FROM purchase_orders po
			WHERE id = $1`

	var po domain.PurchaseOrder

	rows, err := tx.QueryContext(ctx, query, poId)
	helper.PanicError(err)

	if rows.Next() {
		err := rows.Scan(
			&po.Id,
			&po.FactoryName,
			&po.PICName,
			&po.QuantityPO,
			&po.QuantityProduction,
			&po.Item,
			&po.PaymentTerm,
			&po.CreatedAt,
			&po.ExpiredAt,
			&po.UnitItem,
			&po.Description,
			&po.Note,
			&po.Status,
			&po.StatusHistory,
			&po.PoCodeId,
		)
		helper.PanicError(err)
		return po, nil
	} else {
		return po, errors.New("PO is not found")
	}
}

func (repository *PurchaseOrder) SavePurchaseOrder(ctx context.Context, tx *sql.Tx, po domain.PurchaseOrder) (domain.PurchaseOrder, error) {
	query := `INSERT INTO purchase_orders(
			factory_name,
			pic_name,
			quantity_po,
			quantity_production,
			item,
			payment_term,
			created_at,
			expired_at,
			unit_item,
			description,
			note,
			status,
			status_history,
			po_code_id)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)
		RETURNING id`

	err := tx.QueryRowContext(ctx, query,
		&po.FactoryName,
		&po.PICName,
		&po.QuantityPO,
		&po.QuantityProduction,
		&po.Item,
		&po.PaymentTerm,
		&po.CreatedAt,
		&po.ExpiredAt,
		&po.UnitItem,
		&po.Description,
		&po.Note,
		&po.Status,
		&po.StatusHistory,
		&po.PoCodeId).Scan(&po.Id)
	if err != nil {
		return po, err
	}

	return po, nil
}

func (repository *PurchaseOrder) UpdatePurchaseOrder(ctx context.Context, tx *sql.Tx, po domain.PurchaseOrder, poIds int64) (domain.PurchaseOrder, error) {
	query := `UPDATE purchase_orders
		SET
			factory_name=$1,
			pic_name=$2,
			quantity_po=$3,
			quantity_production=$4,
			item=$5,
			payment_term=$6,
			created_at=$7,
			expired_at=$8,
			unit_item=$9,
			description=$10,
			note=$11,
			status=$12,
			status_history=$13,
			po_code_id=$14`

	res, err := tx.ExecContext(ctx, query,
		&po.FactoryName,
		&po.PICName,
		&po.QuantityPO,
		&po.QuantityProduction,
		&po.Item,
		&po.PaymentTerm,
		&po.CreatedAt,
		&po.ExpiredAt,
		&po.UnitItem,
		&po.Description,
		&po.Note,
		&po.Status,
		&po.StatusHistory,
		&po.PoCodeId)
	helper.PanicOnErrorContext(ctx, err)
	r, err := res.RowsAffected()
	helper.PanicOnErrorContext(ctx, err)
	if r == 0 {
		return po, exceptioncode.ErrEmptyResult
	}

	return po, nil
}

func (repository *PurchaseOrder) DeletePurchaseOrder(ctx context.Context, tx *sql.Tx, poId int64) {
	sql := "DELETE FROM purchase_orders WHERE id = $1"

	_, err := tx.ExecContext(ctx, sql, poId)
	helper.PanicOnErrorContext(ctx, err)
}

func (repository *PurchaseOrder) FindPoCode(ctx context.Context, tx *sql.Tx, codeId int) (domain.POCode, error) {
	query := `SELECT
				id,
				code,
				name
			FROM po_codes pc
			WHERE code = $1`

	var pc domain.POCode

	rows, err := tx.QueryContext(ctx, query, codeId)
	helper.PanicError(err)

	if rows.Next() {
		err := rows.Scan(
			&pc.Id,
			&pc.CodeId,
			&pc.Name,
		)
		helper.PanicError(err)
		return pc, nil
	} else {
		return pc, errors.New("PO is not found")
	}
}
