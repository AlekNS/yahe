package services

import (
	"fmt"

	"github.com/alekns/yahe/internal/auth/app"
	"github.com/alekns/yahe/internal/auth/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type jwtServiceImpl struct {
	settings *config.JwtSettings
}

// Sign .
func (s *jwtServiceImpl) SignAndEncode(attrs map[string]interface{}) (string, error) {
	if attrs == nil {
		return "", errors.New("failed to SignAndEncode, attrs is nil")
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(attrs))

	return token.SignedString(s.settings.Secret)
}

// Verify .
func (s *jwtServiceImpl) VerifyAndDecode(tokenString string) (map[string]interface{}, error) {
	if len(tokenString) == 0 {
		return nil, app.ErrorInvalidJwtToken
	}

	var token *jwt.Token
	var err error

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// s.settings.Algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.settings.Secret, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt decode failed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, app.ErrorInvalidJwtToken
}

// NewJwtServiceImpl .
func NewJwtServiceImpl(settings *config.Settings) app.JwtService {
	return &jwtServiceImpl{
		settings: settings.Jwt,
	}
}
