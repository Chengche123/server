package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type jwtVerifier struct {
	publicKey *rsa.PublicKey
}

// Verify Verify the token and return accountID, or return error
func (j *jwtVerifier) VerifyToken(token string) (string, error) {
	tkn, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannt parse token: %v", err)
	}

	if !tkn.Valid {
		return "", fmt.Errorf("token is invalid")
	}

	claims, ok := tkn.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("wront claims")
	}

	if err := claims.Valid(); err != nil {
		return "", fmt.Errorf("invalid claims: %v", err)
	}

	return claims.Subject, nil
}

func NewJWTVerifier(publicKey string) (*jwtVerifier, error) {
	pkey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return nil, fmt.Errorf("invalid publickey: %v", err)
	}

	return &jwtVerifier{
		publicKey: pkey,
	}, nil
}
