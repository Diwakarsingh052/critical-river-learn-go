package main

import (
	"database/sql"
	"log/slog"
	"user-service/handlers"
	"user-service/internal/stores/postgres"
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		slog.Error("error in loading env file")
		return
	}
	// setting up required things for the project
	setup()

}

func setup() {
	/*
			//------------------------------------------------------//
		                Setting up DB & Migrating tables
			//------------------------------------------------------//
	*/

	db, err := setupDatabase()
	if err != nil {
		panic(err)
	}

	conf, err := users.NewConf(db)
	if err != nil {
		panic(err)
	}

	// setting up routes, and running the server
	r := gin.New()
	handlers.RegisterRoutes(r, conf)
	err = r.Run(":80")
	if err != nil {
		panic(err)
	}
}

func setupDatabase() (*sql.DB, error) {
	slog.Info("Connecting to postgres")
	postgresDb, err := postgres.OpenDb()
	if err != nil {
		return nil, err
	}

	slog.Info("Migrating tables for user-service if not already done")
	err = postgres.RunMigrations(postgresDb)
	if err != nil {
		return nil, err
	}

	return postgresDb, nil
}
