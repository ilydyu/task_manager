package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(ID int64, secret string) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    "token_manager",
		Subject:   fmt.Sprintf("%d", ID),
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString, secret string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
