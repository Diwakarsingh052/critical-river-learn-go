package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("sending notification to %s %s %s\n", u.name, u.email, p)
	return len(p), nil
}

func main() {
	var u user = user{
		name:  "james",
		email: "james@email.com",
	}
	// we are able to pass user struct to log.New
	// because user implements io.Writer interface
	l := log.New(u, "sales: ", log.LstdFlags|log.Lshortfile)
	l.Println("hello")
}
