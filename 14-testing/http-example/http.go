package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// http:localhost:8080/double?v=2
// double handler would return 4 (double of 2)
func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("v")
	if text == "" {
		http.Error(w, "missing query parameter", http.StatusBadRequest)
		return
	}
	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "double of number:", v, ":", v*2)

}
