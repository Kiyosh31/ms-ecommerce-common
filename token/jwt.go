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
	email string,
	role string,
	duration int,
) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Duration(duration) * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(c.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	return ss, nil
}

func (c *JwtCreator) VerifyToken(tokenString string) (claims jwt.MapClaims, err error) {
	// Use the jwt.Parse function to parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method (HS256 in this case)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		// Return the secret key for signing verification
		return []byte(c.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract claims from the valid token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unable to convert claims to MapClaims")
	}

	return claims, nil
}
