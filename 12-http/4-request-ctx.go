package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/text", data)
	http.ListenAndServe(":8080", nil)
}

func data(w http.ResponseWriter, r *http.Request) {

	time.Sleep(5 * time.Second)
	// taking the context from the request
	// don't create a new context with context.Background()
	ctx := r.Context()

	// check if the context is canceled or timeout
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println("user cancelled the request")
		return
	default:
		// client is still there
	}
	log.Println("resp sent to the client")
	w.Write([]byte("Hello World"))
}
