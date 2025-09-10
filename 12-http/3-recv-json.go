package main

import (
	"encoding/json"
	"io"
)

func main() {

	var u user
	// r *http.Request
	data, err := io.ReadAll(r.Body)
	err := json.Unmarshal(data, &u)

	// create a handler function that receives json request
	// and covert that into a struct

	// run http server and register one endpoint that receives json
	// vscode extension
	//Thunder Client
}
