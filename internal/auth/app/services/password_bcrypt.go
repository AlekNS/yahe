package services

import (
	"github.com/alekns/yahe/internal/auth/app"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/unicode/norm"
)

type passwordServiceBcrypt struct {
	cost int
}

// Create .
func (ps *passwordServiceBcrypt) Create(salt []byte, passwordStr string) ([]byte, error) {
	var password = norm.NFC.Bytes([]byte(passwordStr))
	password = append(password, salt...)
	return bcrypt.GenerateFromPassword(password, ps.cost)
}

// Verify .
func (ps *passwordServiceBcrypt) Verify(salt []byte, srcPasswordStr string, comparedPassword []byte) (bool, error) {
	var passwordSrc = norm.NFC.Bytes([]byte(srcPasswordStr))
	passwordSrc = append(passwordSrc, salt...)
	var result = bcrypt.CompareHashAndPassword(comparedPassword, passwordSrc)
	if result == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	return result == nil, result
}

// NewPasswordServiceBcrypt .
func NewPasswordServiceBcrypt(cost int) app.PasswordService {
	return &passwordServiceBcrypt{
		cost: cost,
	}
}
