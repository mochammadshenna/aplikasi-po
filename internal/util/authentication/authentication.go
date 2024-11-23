package authentication

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mochammadshenna/aplikasi-po/internal/util/exceptioncode"
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
)

type (
	key     string
	Payload struct {
		IDUser int64 `json:"id_user"`
		jwt.StandardClaims
	}
)

const JWTClaim key = "jwtclaim"

var (
	secretKey string
)

func Init(key string) {
	if len(key) < 32 {
		panic(errors.New("invalid key size"))
	}
	secretKey = key
}

func CreateToken(duration time.Duration, idUser int64) string {
	tokenPayload := Payload{
		IDUser: idUser,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenPayload)
	token, err := jwtToken.SignedString([]byte(secretKey))
	helper.PanicError(err)
	return token
}

func VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, exceptioncode.ErrTokenInvalid
		}
		return []byte(secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("token has expired")) {
			return nil, exceptioncode.ErrTokenExpired
		}
		if strings.Contains(verr.Error(), "token is expired") {
			return nil, exceptioncode.ErrTokenExpired
		}
		return nil, exceptioncode.ErrTokenInvalid
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, exceptioncode.ErrTokenInvalid
	}

	return payload, nil
}

func ExtractClaim(request *http.Request) *Payload {
	jwtClaim := request.Context().Value(JWTClaim)
	if jwtClaim == nil {
		return &Payload{}
	}
	return jwtClaim.(*Payload)
}

func ExtractClaimRequest(r *http.Request) *Payload {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return &Payload{}
	}
	if !strings.Contains(reqToken, "Bearer ") {
		return &Payload{}
	}
	s := strings.Split(reqToken, " ")
	claims, err := VerifyToken(s[1])
	if err != nil {
		return &Payload{}
	}
	return claims
}
