package main

import (
	"encoding/json"
	"log"
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

	// NewEncoder will encode the struct to json and write the response to the client

	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errMap := map[string]any{"error": "error in encoding"}
		json.NewEncoder(w).Encode(errMap)
	}
	log.Println("resp sent to the client")
	// marshal converts the type to json
	//jsonData, err := json.Marshal(u)
	//if err != nil {
	//	// text based error
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//
	// write the response to the client
	//w.Write(jsonData)
}
