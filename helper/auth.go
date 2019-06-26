package helper

import (
	"context"
	"errors"
	"firebase.google.com/go/auth"
	"github.com/ToluwaniO/utube/global"
	"strings"
)

func ValidateToken(ctx context.Context, authorization string) (*auth.Token, error)  {
	idToken := strings.Split(authorization, " ")[1]
	token, err := global.Auth.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, errors.New("could not validate id token")
	}
	return token, nil
}

func GetUserToken(ctx context.Context, idToken string) *auth.Token {
	token, err := ValidateToken(ctx, idToken)

	if err != nil {
		return nil
	}
	return token
}
