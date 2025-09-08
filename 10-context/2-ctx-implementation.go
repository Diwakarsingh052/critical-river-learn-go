package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	//sql.DB.ExecContext()
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()
	// slow function would return an error if the context is canceled or timeout happens
	// it also returns an error if the string conversion fails
	i, err := Slow(ctx, "10")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("slow func returned the result", i)
}

// Context must be passed as the first argument

func Slow(ctx context.Context, msg string) (int, error) {
	//sql.DB{}.ExecContext()
	i, err := strconv.Atoi(msg)
	time.Sleep(time.Second * 2)
	if err != nil {
		return 0, err
	}
	select {
	case <-ctx.Done():
		// this case would be true if the context is timer is over or canceled
		return 0, ctx.Err()
	default:
		// time is still there
		// if there use case:
		// rollback your previous work if you don't want to persist the changes
	}
	return i, nil
}
