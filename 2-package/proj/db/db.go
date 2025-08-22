package db

import "fmt"

// var DbConn = "mysql"
var dbConn string

func IntializeDb(conn string) {
	dbConn = conn
}

func GetConn() {
	fmt.Println(dbConn)
}
