package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

//	"github.com/jackc/pgx/v5/pgxpool"
//
// pgxpool supports connection pooling, useful for concurrent access to the database.
func main() {

	db := ConnectToDB()
	defer db.Close()
	//CreateTable(db)
	//insertUser(db, "Bob", "bob@example.com", 35)
	//updateUserEmail(db, 1, "abc@email.com")
	getAllUsers(db)
}
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

// Three methods to execute queries on the database
// Exec -> when query does not return anything
// QueryRow -> returns exactly one row
// Query -> returns many rows

func CreateTable(db *pgxpool.Pool) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100),
        email VARCHAR(100) UNIQUE NOT NULL,
        age INT
    );`

	res, err := db.Exec(context.Background(), query)
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows affected: %d\n", res.RowsAffected())

}

func insertUser(db *pgxpool.Pool, name, email string, age int) {
	// SQL query to insert a user and return the new user's ID
	// don't hardcode the values, or use the string in construction, sql injection can happen
	q := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id`
	ctx := context.Background()
	var id int
	// Execute the query to insert the user and get the new user's ID
	//QueryRow returns one row as output
	err := db.QueryRow(ctx, q, name, email, age).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User inserted successfully", id)

}

// updateUserEmail updates a user's email based on their ID.
func updateUserEmail(db *pgxpool.Pool, userID int, newEmail string) {
	// SQL query to update a user's email
	query := `UPDATE users SET email = $1 WHERE id = $2`

	// prone to SQL injection, we should not construct queries using strings
	//query := fmt.Sprintf("SELECT * FROM users WHERE username = '%s'", "userInput")

	// Execute the query to update the user's email
	_, err := db.Exec(context.Background(), query, newEmail, userID)
	if err != nil {
		log.Fatalf("Unable to update user: %v\n", err) // Log and terminate if update fails
	}
	fmt.Println("User email updated")
}

func getAllUsers(db *pgxpool.Pool) {
	// user query method
	// run a for loop on rows.Next()
	// inside the loop scan values using rows.Scan

	type User struct {
		ID    int
		Name  string
		Email string
		Age   int
	}
	// SQL query to retrieve all users
	query := `SELECT id, name, email, age FROM users`

	rows, err := db.Query(context.Background(), query)
	if err != nil {
		// log.Fatal should be avoided in production code
		// it quits the program immediately
		// only in startup code it should be fine
		log.Fatalf("Unable to retrieve users: %v\n", err)
	}
	defer rows.Close()

	var users []User
	// this loop would run until there are rows to scan
	for rows.Next() {
		var u User

		// Scan the values from the row, and store them in the user struct
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			log.Printf("Unable to scan row: %v\n", err)
			continue
		}
		users = append(users, u)
	}

	fmt.Printf("Users: %+v\n", users)

}
