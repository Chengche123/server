package token

import (
	"errors"
	"time"

	xerrors "github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

var (
	ErrInvalidKey = errors.New("auth-service/token: invalid key")
	ErrSigned     = errors.New("auth-service/token: failed to signed")
)

type JWTTokenGen struct {
	privateKey []byte
	iss        string

	nowFunc func() time.Time
}

func NewJWTTokenGen(privateKey []byte, iss string) (*JWTTokenGen, error) {
	return &JWTTokenGen{
		privateKey: privateKey,
		iss:        iss,
		nowFunc:    time.Now,
	}, nil
}

func (j *JWTTokenGen) GenerateToken(accountID string, expireIn time.Duration) (string, error) {
	now := j.nowFunc().Unix()

	text := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		ExpiresAt: now + int64(expireIn.Seconds()),
		IssuedAt:  now,
		Issuer:    j.iss,
		Subject:   accountID,
	})

	pkey, err := jwt.ParseRSAPrivateKeyFromPEM(j.privateKey)
	if err != nil {
		return "", xerrors.Wrapf(ErrInvalidKey, "jwt: [%v]", err)
	}

	token, err := text.SignedString(pkey)
	if err != nil {
		return "", xerrors.Wrapf(ErrSigned, "jwt: [%v]", err)
	}

	return token, err
}
