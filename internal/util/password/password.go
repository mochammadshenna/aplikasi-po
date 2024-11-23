package password

import (
<<<<<<< HEAD
	"github.com/mochammadshenna/aplikasi-po/internal/util/helper"
	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	helper.PanicError(err)
	return string(bytes)
}

func CheckHashPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
=======
	"golang.org/x/crypto/bcrypt"
)

// HashPassword creates a bcrypt hash from password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckHashPassword compares a bcrypt hashed password with its possible plaintext equivalent
func CheckHashPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
>>>>>>> ffd4b1225fa304d1a73819bffb534cf23222fb2f
}
