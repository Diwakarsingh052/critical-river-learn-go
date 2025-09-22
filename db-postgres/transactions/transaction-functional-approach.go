package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	db := ConnectToDBV2()
	defer db.Close()
	fmt.Println(Update(db))
}

func ConnectToDBV2() *pgxpool.Pool {
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

func Update(db *pgxpool.Pool) error {
	f := func(tx pgx.Tx) error {
		//
		updateQuery := `UPDATE author
					SET name = $1
					WHERE email = $2;`

		res, err := tx.Exec(context.Background(), updateQuery, "abc", "john.doe@example.com")
		if err != nil {
			return err
		}
		if res.RowsAffected() == 0 {
			return fmt.Errorf("no rows affected")
		}

		res, err = tx.Exec(context.Background(), updateQuery, "xyz", "jane.smith@example.com")
		if err != nil {
			return err
		}
		if res.RowsAffected() == 0 {
			return fmt.Errorf("no rows affected")
		}
		return nil
	}

	err := withTransaction(db, f)
	if err != nil {
		return err
	}
	return nil
}

func withTransaction(db *pgxpool.Pool, fn func(tx pgx.Tx) error) error {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		err := tx.Rollback(ctx)
		if err != nil {
			return fmt.Errorf("unable to rollback transaction: %w", err)
		}
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}
	return nil
}
