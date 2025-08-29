package main

import (
	"fmt"
	"learn-go/7-structs/database"
)

func main() {
	c, err := database.NewConn("mysql")
	if err != nil {
		panic(err)
	}
	c.AddUser()
	fmt.Println(c)

}
