package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func GenerateCredentialsHash(credentials string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(credentials), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), err
}

func ValidateCredentials(credentials, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(credentials))
}