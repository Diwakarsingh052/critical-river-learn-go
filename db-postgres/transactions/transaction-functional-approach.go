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
	// Business logic function, which takes a transaction object as input
	// and performs database operations within the transaction
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

	// Execute the business logic function within a transaction
	// withTransaction handles all the transaction lifecycle (begin, commit, rollback)
	err := withTransaction(db, f)
	if err != nil {
		return err
	}
	return nil
}

// withTransaction  function that provides transaction management.
// It takes a database pool and a function that contains the business logic,
// then handles the transaction lifecycle automatically.

func withTransaction(db *pgxpool.Pool, fn func(tx pgx.Tx) error) error {
	ctx := context.Background()

	// Begin a new database transaction
	tx, err := db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	// Execute the business logic function
	// Pass the transaction object so it can perform database operations
	err = fn(tx)
	if err != nil {
		errRollback := tx.Rollback(ctx)
		if errRollback != nil {
			return fmt.Errorf("unable to rollback transaction: %w", err)
		}
		// If business logic failed, rollback the transaction
		return err
	}

	// If business logic succeeded, commit the transaction
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("unable to commit transaction: %w", err)
	}
	return nil
}

/*
┌─────────────────┐
│   Update()      │
│                 │
│ 1. Define f()   │
│ 2. Call         │
│    withTransaction│
└─────────┬───────┘
          │
	      ▼
┌─────────────────┐
│ withTransaction()│
│                 │
│ 1. Begin TX     │
│ 2. Call f(tx)   │
│ 3. Commit/      │
│    Rollback     │
└─────────┬───────┘
		  │
		  ▼
┌─────────────────┐
│     f(tx)       │
│                 │
│ 1. UPDATE #1    │
│ 2. Check rows   │
│ 3. UPDATE #2    │
│ 4. Check rows   │
└─────────────────┘
*/
