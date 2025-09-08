package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// context is a container to hold values
	// context can hold timer values or normal values as well
	// we need a container to store context values

	// if no context is available, we will create a new container to store the context values

	// create a new empty context // which is an empty container
	ctx := context.Background()
	// Todo is a context with no values
	// but it gives a signal that we are not sure what context to use , empty or from somewhere else
	//ctx := context.TODO()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	//context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	// if cancel is called without defer, context will be canceled immediately
	defer cancel() // cancel cleanup the resources taken by the context

	// constructing a new request with context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Println(err)
		return
	}

	// making the request to url
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	// reading the response body which would be in bytes
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// printing the response body
	fmt.Println(string(data))

}
