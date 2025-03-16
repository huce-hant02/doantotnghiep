package middlewares

import (
	"context"
	"doantotnghiep/doantotnghiep/infrastructure"
	"errors"

	"github.com/lestrrat-go/jwx/jwt"
)

const (
	TypeRefresh = "TypeRefresh"
	TypeAccess  = "TypeAccess"
)

func VerifyToken(tokenString string, typeToken string) (jwt.Token, error) {
	// Verify the token
	token, err := infrastructure.GetDecodeAuth().Decode(tokenString)
	if err != nil {
		return nil, err
	}
	// Check token valid
	if _, ok := token.AsMap(context.Background()); ok != nil {
		return nil, errors.New("Token invalid")
	}

	var tokenUUID string
	var ok bool
	claims, _ := token.AsMap(context.Background())
	if typeToken == TypeRefresh {
		tokenUUID, ok = claims["refresh_uuid"].(string)
		if !ok {
			return nil, errors.New("Claims of token is invalid")
		}
	} else {
		tokenUUID, ok = claims["access_uuid"].(string)
		if !ok {
			return nil, errors.New("Claims of token is invalid")
		}
	}
	if userID, err := infrastructure.FetchAuth(tokenUUID); err != nil || userID == 0 {
		return nil, errors.New("Token is expired")
	}

	return token, nil
}
