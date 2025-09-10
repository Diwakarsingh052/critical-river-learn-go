package main

import (
	"encoding/json"
	"io"
)

// https://mholt.github.io/json-to-go/
/*
{
    "glossary": {
        "title": "example glossary",
		"GlossDiv": {
            "title": "S",
			"GlossList": {
                "GlossEntry": {
                    "ID": "SGML",
					"SortAs": "SGML",
					"GlossTerm": "Standard Generalized Markup Language",
					"Acronym": "SGML",
					"Abbrev": "ISO 8879:1986",
					"GlossDef": {
                        "para": "A meta-markup language, used to create markup languages such as DocBook.",
						"GlossSeeAlso": ["GML", "XML"]
                    },
					"GlossSee": "markup"
                }
            }
        }
    }
}
*/
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
