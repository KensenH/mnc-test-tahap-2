package auth

import (
	"context"
	"fmt"
	"my-rest-api/internal/domain/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s Auth) Login(ctx context.Context, req auth.LoginReq) (string, error) {
	user, err := s.Data.GetUserByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		return "", err
	}

	if user.Pin != req.Pin {
		return "", fmt.Errorf("phone Number and PIN doesn't match.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.UserID,
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, err
}
