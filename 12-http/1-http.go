package main

import (
	"fmt"
	"net/http"
	"sync"
)

// curl localhost:8080/home // to query the server
// go run main.go // to run the server

// by default http service runs all the requests in separate goroutines
// if panic happens in a request, http can automatically recover the panic and send the error to the client

// but if panic happens in a goroutine which we have manually created, http can't recover the panic'
// and the service will crash

func main() {
	// HandleFunc register the route and the function that would handle the request
	// HandleFunc takes function as argument for handling the request
	http.HandleFunc("/home", Home)
	// http.ListenAndServe starts a server on :8080
	// ListenAndServe would run forever
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct

	// add recovery to the goroutine to stop the crashing of the server
	wg := new(sync.WaitGroup)
	wg.Go(func() {
		fmt.Println("doing some internal jobs when home is called")
		panic("some error")
	})
	wg.Wait()
	//panic("something went wrong")
	w.Write([]byte("hello"))
}
