package helper

import (
	"context"
	"database/sql"

	"github.com/lib/pq"
	"github.com/mochammadshenna/aplikasi-po/internal/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/internal/util/logger"
)

func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicOnErrorContext(ctx context.Context, err error) {
	if err != nil {
		logger.Error(ctx, err)
		panic(err)
	}
}

func TranslatePostgreError(ctx context.Context, err error) error {
	if err != nil {
		if err == sql.ErrNoRows {
			err = exceptioncode.ErrEmptyResult
		} else if pgErr, isPGErr := err.(*pq.Error); isPGErr {
			if pgErr.Code == "23503" {
				err = exceptioncode.ErrForeignKeyViolation
			} else if pgErr.Code == "23505" {
				err = exceptioncode.ErrUniqueViolation
			} else {
				logger.Error(ctx, err)
			}
		} else {
			logger.Error(ctx, err)
		}
	}
	return err
}
