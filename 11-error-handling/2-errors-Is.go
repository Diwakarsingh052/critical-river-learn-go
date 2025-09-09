package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var userList = make(map[int]string)
var ErrNotFoundV2 = errors.New("not found")

// error must be the last value to be returned from function
// whenever error happens, set other values to default

func FetchRecordV2(idString string) (string, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return "", err
	}
	name, ok := userList[id]
	if !ok {
		return "", ErrNotFoundV2
		//return "", errors.New("not found")
		//return "", fmt.Errorf("not found %s", "some msg")
	}
	return name, nil
}

func main() {
	// default value of error is nil
	name, err := FetchRecordV2("101")
	if err != nil {
		// errors.Is check if err is present in the error chain
		// if it is present, then we can take specific action to handle it
		// in this case we are looking for ErrNotFoundV2 inside the error chain
		if errors.Is(err, ErrNotFoundV2) {
			fmt.Println("sending a customized msg to the end user")
			fmt.Println("create your account first")
			log.Println(err)
			return
		}

		log.Println(err)
		return

	}

	fmt.Println(name)
}
