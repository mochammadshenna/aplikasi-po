package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/mochammadshenna/aplikasi-po/internal/util/authentication"
)

// JWTAuth get token from header "Authorization" with bearer format
// and pass extracted token to context
func JWTAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !strings.Contains(reqToken, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		s := strings.Split(reqToken, " ")
		claims, err := authentication.VerifyToken(s[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), authentication.JWTClaim, claims)
		h(w, r.WithContext(ctx), ps)
	}
}
