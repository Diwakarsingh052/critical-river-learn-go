package main

import (
	"fmt"
	"math/rand"
	"proj/db"
	"proj/sum"
) // moduleName/packageName

// package name must be in lower case, no -, _, no numbers

// create package names according to the work they do

func main() {
	fmt.Println(sum.Add(10, 20))
	fmt.Println(sum.Sum(10, 2))
	//fmt.Println(db.DbConn)
	//db.DbConn = "pg"
	db.IntializeDb("mysql")
	sum.Insert()
	db.GetConn()
	//ioutil.ReadAll()
	//io.ReadAll()
	//ioutil.ReadFile()
	//os.ReadFile()
	rand.Int()
}
