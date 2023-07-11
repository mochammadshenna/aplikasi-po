package helper

import (
	"context"

	"github.com/mochammadshenna/aplikasi-po/util/logger"
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
