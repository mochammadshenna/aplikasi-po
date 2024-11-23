package password

import (
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
}
