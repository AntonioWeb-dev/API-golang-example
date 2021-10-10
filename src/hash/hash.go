package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash - receive a password as a param and return a password hashed
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func Verify(password string, passwordHash string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
