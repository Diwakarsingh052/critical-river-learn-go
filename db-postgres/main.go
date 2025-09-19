package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//	"github.com/jackc/pgx/v5/pgxpool"
//
// pgxpool supports connection pooling, useful for concurrent access to the database.
func main() {

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

	config, err := pgxpool.ParseConfig(psqlInfo)
	if err != nil {
		panic(err)
	}
	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = 60 * time.Minute
	config.HealthCheckPeriod = 10 * time.Second

	// Connect to the database
	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Ping the database to check if it is alive
	err = db.Ping(context.Background())
	if err != nil {
		panic(err)
	}
}
