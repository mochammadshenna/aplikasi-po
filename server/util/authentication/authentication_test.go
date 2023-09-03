package authentication_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/mochammadshenna/aplikasi-po/util/authentication"
	"github.com/stretchr/testify/assert"
)


func TestAuthentication(t *testing.T) {
	authentication.Init("ylEHnM9TGlVtNHFnpcmxp6c9BGJ6XqZX")
	var idUser int64 = 10
	token := authentication.CreateToken(time.Minute, idUser)
	claim, err := authentication.VerifyToken(token)
	assert.NoError(t, err)
	request := http.Request{}
	ctx := context.WithValue(request.Context(), authentication.JWTClaim, claim)
	r := request.WithContext(ctx)
	payload := authentication.ExtractClaim(r)
	assert.Equal(t, idUser, payload.IDUser)
}