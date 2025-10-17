package main

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLXNlcnZpY2UiLCJzdWIiOiIxMDEiLCJleHAiOjE3NjA3MDAzNzIsImlhdCI6MTc2MDY5NzM3Miwicm9sZXMiOlsiYWRtaW4iLCJ1c2VyIl19.eDF1TE8arBC2CM5hKrc1Tey8PaL_O2CTieVj87iZjwrISbq5qVdQGigj21tr3ZHBYf5wijA3aOq7E0YFWsk2PmyT3k4l6XjldXHz9US1ayhm9iZ4NfCvmqiJTTfx8qUedCo5ggDQhegxyBeR6J_iPq4Dqk4LAZ3AfYGPePQj-5th1xLlm_23TAO2f3rGrZosM8aMtU-u8EldsTII7RzfYzY2svNiYnp9hBOoAFS-a8ZlD8oGrqi9l2pUlkIYdNqtjVq7Qjqq3m7l4IkVjbQwLyVk1YfOR5eVmnBQ9j3ICdtqgwhGKWccvXrXGUo5EPPpEi8GZ3JGFUqpOaBCzLk5uw`

func main() {

	f, err := os.ReadFile("pubkey.pem")

	if err != nil {
		panic(err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(f)

	if err != nil {
		panic(err)
	}
	var claims struct {
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}
	token, err := jwt.ParseWithClaims(tokenStr, &claims, func(token *jwt.Token) (any, error) {
		return publicKey, nil
	})
	if err != nil {
		panic(err)
	}
	if !token.Valid {
		panic("token is invalid")
	}
	fmt.Printf("%+v", claims)

}
