package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

func main() {

	privateKeyPem, err := os.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyPem)
	if err != nil {
		panic(err)
	}
	// iss (issuer): Issuer of the JWT
	// sub (subject): Subject of the JWT (the users)
	// aud (audience): Recipient for which the JWT is intended
	// exp (expiration time): Time after which the JWT expires
	// nbf (not before time): Time before which the JWT must not be accepted for processing
	// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
	// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)
	c := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "user-service",
			Subject:   "101",                                                // userId
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(50 * time.Minute)), // after 50 minutes, this token expires
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Roles: []string{"admin", "user"},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	signedToken, err := tkn.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(signedToken)
}
