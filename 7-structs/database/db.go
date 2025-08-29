package database

import (
	"errors"
	"fmt"
)

type Conn struct {
	//db       *sql.DB
	dbString string
}

func NewConn(dbString string) (*Conn, error) {
	if dbString == "" {
		return nil, errors.New("invalid db or dbString")
	}
	c := &Conn{dbString: dbString}
	return c, nil
}

func (c *Conn) AddUser() {
	fmt.Println("adding user", c.dbString)
}
