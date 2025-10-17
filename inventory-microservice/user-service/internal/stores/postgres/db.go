package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	driverName    = "pgx" // Database driver
	migrationsDir = "internal/stores/postgres/migrations"
)

func OpenDb() (*sql.DB, error) {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port     = os.Getenv("POSTGRES_PORT")
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DATABASE")
	)

	//sql.Open(psqlInfo)
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var backoff time.Duration = 2
	var err error
	var db *sql.DB

	for i := 0; i < 8; i++ {
		db, err = sql.Open(driverName, connectionString)
		if err != nil {
			fmt.Println("postgres not ready yet")
			time.Sleep(backoff * time.Second)
			backoff += 2
			continue
		}
		err = db.Ping()
		if err != nil {
			fmt.Println("postgres not ready yet")
			time.Sleep(backoff * time.Second)
			backoff += 2
			continue
		}
		break
	}

	if err != nil {
		return nil, err
	}
	return db, nil

}

func RunMigrations(db *sql.DB) error {

	// Set the dialect for Goose (PostgreSQL in this case)
	err := goose.SetDialect(driverName)
	if err != nil {
		return err
	}

	// Apply all pending migrations in the directory
	err = goose.Up(db, migrationsDir)
	if err != nil {
		return err
	}
	return nil
}
