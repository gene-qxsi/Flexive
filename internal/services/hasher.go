package services

import "golang.org/x/crypto/bcrypt"

type PasswordHasher interface {
	Hash(password string) string
}

type BcryptHasher struct {
	salt string
}

func NewSHA1Hasher(salt string) *BcryptHasher {
	return &BcryptHasher{salt: salt}
}

func (h *BcryptHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (h *BcryptHasher) Compare(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
