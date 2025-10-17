package auth

import (
	"crypto/rsa"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}
type Keys struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewKeys(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Keys, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("private key or public key is nil")
	}
	return &Keys{privateKey: privateKey, publicKey: publicKey}, nil
}

func (k *Keys) GenerateToken(claims Claims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := tkn.SignedString(k.privateKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (k *Keys) ValidateToken(tokenStr string) (Claims, error) {
	var claims Claims
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return k.publicKey, nil
	})
	if err != nil {
		return Claims{}, err
	}

	if !tkn.Valid {
		return Claims{}, errors.New("token is invalid")
	}
	return claims, nil
}
