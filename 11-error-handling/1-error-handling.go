package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

var user = make(map[int]string)
var ErrNotFound = errors.New("not found")

// error must be the last value to be returned from function
// whenever error happens, set other values to default

func FetchRecord(idString string) (string, error) {
	id, err := strconv.Atoi(idString)
	if err != nil {
		return "", err
	}
	name, ok := user[id]
	if !ok {
		return "", ErrNotFound
		//return "", errors.New("not found")
		//return "", fmt.Errorf("not found %s", "some msg")
	}
	return name, nil
}

func main() {
	// default value of error is nil
	name, err := FetchRecord("abc")
	if err != nil {
		log.Println(err)
		return
		// structured logging
		// slog
		// zerolog, logrus, zap, etc.

		// avoid using Fatal, or os.Exit
		// because it would kill the program immediately
	}

	fmt.Println(name)
}
