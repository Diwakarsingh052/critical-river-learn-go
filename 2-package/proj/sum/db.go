package sum

import (
	"fmt"
	"proj/db"
)

func Insert() {
	db.IntializeDb("pg")
	fmt.Println("Insert")
}
