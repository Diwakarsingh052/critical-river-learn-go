package mysql

import (
	"app/internal/stores/models"
	"fmt"
)

type Conn struct {
	userDb map[int]models.User
}

func NewConn() *Conn {
	return &Conn{
		userDb: make(map[int]models.User),
	}
}

func (c *Conn) Create(u models.User) (models.User, error) {

	fmt.Println("Creating a user in mysql", " u : ", u)
	//Need to check if user exists if yes then throw error else save
	c.userDb[u.Id] = u // adding key value pair

	return u, nil

}

func (c *Conn) Delete(id int) bool {
	u, ok := c.userDb[id] // check if user exists

	if !ok { // ok == false
		fmt.Println("User with id ", id, "Is not found for delete")
		return false
	}
	fmt.Println("deleting a user in map mysql", " u : ", u)

	delete(c.userDb, id)
	return true
}
