package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtCreator struct {
	secretKey string
}

func NewJwtCreator(secretKey string) *JwtCreator {
	return &JwtCreator{
		secretKey: secretKey,
	}
}

func (creator *JwtCreator) CreateToken(
	id string,
	email string,
	isAdmin bool,
	duration time.Duration,
) (
	string,
	*UserClaims,
	error,
) {
	claims, err := NewUserClaims(id, email, isAdmin, &duration)
	if err != nil {
		return "", nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(creator.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("error signing token: %v", err)
	}

	return tokenStr, claims, nil
}

func (creator *JwtCreator) VerifyToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// verify with signed method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid signin method")
		}

		return []byte(creator.secretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
