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

func (c *JwtCreator) CreateToken(
	id string,
	email string,
	role string,
	duration time.Duration,
) (
	string,
	*UserClaims,
	error,
) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 1).Unix(), // Expires after 1 hour
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		return "", nil, fmt.Errorf("failed to create token: %w", err)
	}
	return ss, nil, nil

	// claims, err := NewUserClaims(id, email, isAdmin, &duration)
	// if err != nil {
	// 	return "", nil, err
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenStr, err := token.SignedString([]byte(creator.secretKey))
	// if err != nil {
	// 	return "", nil, fmt.Errorf("error signing token: %v", err)
	// }

	// return tokenStr, claims, nil
}

func (c *JwtCreator) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(c.secretKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("token claims are not of type MapClaims")
	}

	return claims, nil
	// token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
	// 	// verify with signed method
	// 	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	// 	if !ok {
	// 		return nil, fmt.Errorf("invalid signin method")
	// 	}

	// 	return []byte(creator.secretKey), nil
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("error parsing token: %v", err)
	// }

	// claims, ok := token.Claims.(*UserClaims)
	// if !ok {
	// 	return nil, fmt.Errorf("invalid token claims")
	// }

	// return claims, nil
}
