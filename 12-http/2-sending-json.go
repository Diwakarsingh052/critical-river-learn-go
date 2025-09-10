package main

import (
	"encoding/json"
	"net/http"
)

// Fields must be exported for json to work
type user struct {
	// use field tags to specify the field name or ignore fields
	FirstName string `json:"first_name"`
	Password  string `json:"-"` // - ignore the field in the output
	Email     string `json:"email"`
}

type mapUser map[string]any

func main() {

	http.HandleFunc("/json", sendJson)
	panic(http.ListenAndServe(":8080", nil))
}

func sendJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//
	u := user{
		FirstName: "abc",
		Password:  "xyz",
		Email:     "abc@gmail.com",
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		// text based error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
