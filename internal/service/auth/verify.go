package auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func (s Auth) VerifyToken(ctx context.Context, tokenString string) (string, error) {
	var (
		err error
		id  string
	)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(s.Secret), nil
	})
	if err != nil {
		return id, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims["id"].(string), err
	} else {
		return "", fmt.Errorf("not valid")
	}
}
