package middleware

import (
<<<<<<< HEAD
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
=======
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from header
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing authorization token",
			})
		}

		// Verify token
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil // Use your secret key from config
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Store user info in context
		c.Locals("user", claims)

		return c.Next()
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
	}
}
