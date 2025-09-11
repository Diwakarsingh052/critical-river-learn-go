package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// go get github.com/google/uuid
// install the external modules

// all the modules are downloaded in the GOPATH folder
// go env GOPATH

// The provided key must be comparable and should not be of type string or
// any other built-in type to avoid collisions between packages using context.
// Users of WithValue should define their own types for keys.
type ctxKey string

const reqIdKey ctxKey = "reqId"

// auth -> key string auth.ctxKey

// reqId -> key string log.ctxKey

// req -> mid1 -> mid2 -> handler -> services -> services -> services

// context values must be used for request scoped data
// something that starts with request and die with request
// like reqId, authToken, jwtToken, etc

func main() {

	http.HandleFunc("/hello", ReqIdMid(Hello))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
func Hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uuid.NewString()
	// getting the reqId from the context
	// checking the type of the value using type assertion
	// ok would be false if the value is not of type string
	reqId, ok := ctx.Value(reqIdKey).(string)
	if !ok {
		reqId = "unknown"
	}
	log.Println(reqId, "sending hello as response")
	fmt.Fprintln(w, "hello")
}
func ReqIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context() // taking the copy of context from the request

		// adding the reqId to the context
		// context stores the key value pair
		ctx = context.WithValue(ctx, reqIdKey, 123)

		fmt.Println("req started with ", reqId)

		// withContext updates the request to store the new context
		next(w, r.WithContext(ctx))
	}
}
