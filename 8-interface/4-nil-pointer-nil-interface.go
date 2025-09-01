package main

import (
	"database/sql"
	"fmt"
)

type DB interface {
	ReadAll()
}

type conn struct {
	db   *sql.DB // default value would be nil
	data string
}

func (c *conn) ReadAll() {
	if c == nil {
		fmt.Println("conn is nil")
		return
	}
	c.data = "some data"
	fmt.Println("read all the values from the db : ", c.data)
	//c.db.Close() // this will panic // db field is nil
}

func main() {
	//c := &conn{} // c is not nil, it will store the address of conn which would be intialized with default values

	var c *conn   // nil pointer
	var db DB = c // / db is not nil // storing a conn type in it
	// if interface is holding any concrete type then it is not nil,
	// even the type value itself could be nil
	// nil interface means no concrete type is present inside the interface

	//var db1 DB // this is nil interface, // it is not holding any concrete type
	if db == nil {
		fmt.Println("db is nil")
		return
	}

	db.ReadAll()
}
