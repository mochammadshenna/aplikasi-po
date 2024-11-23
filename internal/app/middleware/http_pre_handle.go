package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mochammadshenna/aplikasi-po/internal/state"
)

func PreHandle(writer *responseWriter, request *http.Request) (*responseWriter, *http.Request) {
	requestId := uuid.New()

	ctx := context.WithValue(request.Context(), state.HttpHeaders().RequestId, requestId.String())
	ctx = context.WithValue(ctx, state.HttpHeaders().PlatformType, request.Header.Get(state.HttpHeaders().PlatformType.String()))
	ctx = context.WithValue(ctx, state.HttpHeaders().Platform, request.Header.Get(state.HttpHeaders().Platform.String()))
	ctx = context.WithValue(ctx, state.HttpHeaders().Version, request.Header.Get(state.HttpHeaders().Version.String()))
	request = request.WithContext(ctx)

	cacheControl := request.Header.Get(state.HttpHeaders().CacheControl.String())

	writer.Header().Add(state.HttpHeaders().ContentType.String(), state.HttpContentTypeValues().ApplicationJson)
	writer.Header().Add(state.HttpHeaders().StartTime.String(), fmt.Sprintf("%d", time.Now().UnixNano()))
	writer.Header().Add(state.HttpHeaders().RequestId.String(), requestId.String())

	// only return Cache-Control header when the request contains Cache-Control header
	if cacheControl != "" {
		writer.Header().Add(state.HttpHeaders().CacheControl.String(), cacheControl)
	}

	return writer, request
}
