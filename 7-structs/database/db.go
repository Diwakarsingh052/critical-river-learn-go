package database

import (
	"errors"
	"fmt"
)

// Conn struct is exported but fields are unexported, so no one can access or change them
type Conn struct {

	//db       *sql.DB
	dbString string
}

// NewConn would return a new instance of Conn
// because db field is unexported, it can't be accessed outside of the package'
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
