package password

import (
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
}
