package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectToDB() *pgxpool.Pool {
	const (
		host     = "localhost"
		port     = "5434"
		user     = "postgres"
		password = "postgres"
		dbname   = "postgres"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Package pgxpool is a concurrency-safe connection pool for pgx.
	config, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		panic(err)
	}

	// Connection pool configuration
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 60 * time.Minute
	config.HealthCheckPeriod = 10 * time.Second

	// Connect to the database
	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}

	// Ping the database to check if it is alive
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := ConnectToDB()
	defer db.Close()
	err := UpdateAuthor(db)
	if err != nil {
		panic(err)
	}
}

func UpdateAuthor(db *pgxpool.Pool) error {
	tx, err := db.BeginTx(context.Background(), pgx.TxOptions{})
	if err != nil {
		panic(err)
	}
	// calling rollback multiple times have no effect after commit
	// rollback would roll back any changes if function return early without commit
	defer func() {
		err := tx.Rollback(context.Background())
		if err != nil {
			log.Printf("Unable to rollback transaction: %v\n", err)
			return
		}
		fmt.Println("Transaction rolled back")

	}()

	updateQuery := `UPDATE author
					SET name = $1
					WHERE email = $2;`

	res, err := tx.Exec(context.Background(), updateQuery, "john", "john1.doe@example.com")
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	res, err = tx.Exec(context.Background(), updateQuery, "jane", "jane1.smith@example.com")
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("no rows affected")
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return err
	}
	return nil

}
