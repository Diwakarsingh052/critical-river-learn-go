package main

import (
	"fmt"
	"net/http"
)

// middleware
// middleware defines preprocessing or postprocessing logic for a request
// middleware run before the handler

// req -> mid1->mid-2-> handler->services
//
//	<-		<-		<-
func main() {
	http.HandleFunc("/home", Mid1(Mid2(home)))
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home handler")
	fmt.Fprintln(w, "home handler")

}

func Mid1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mid1 layer started")
		fmt.Println("pre processing logic")
		next(w, r) // actual call the handler
		fmt.Println("post processing logic")
		fmt.Println("mid1 layer ended")

	}
}

func Mid2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mid2 layer started")
		fmt.Println("pre processing logic")
		next(w, r) // actual call the handler
		fmt.Println("post processing logic")
		fmt.Println("mid2 layer ended")

	}
}
