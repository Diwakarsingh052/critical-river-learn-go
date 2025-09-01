package main

import (
	"app/internal/stores"
	"app/internal/stores/models"
	"app/internal/stores/mysql"
	"app/internal/stores/postgres"
)

func main() {
	u := models.User{
		Id:   101,
		Name: "Raj",
	}
	m := mysql.NewConn()
	p := postgres.NewConn()
	var ur stores.UserRepository = m

	ur.Create(u)

	ur = p

	ur.Create(u)
}
