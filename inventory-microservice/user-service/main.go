package main

import (
	"database/sql"
	"log/slog"
	"os"
	"user-service/handlers"
	"user-service/internal/auth"
	"user-service/internal/stores/postgres"
	"user-service/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	//------------------------------------------------------//
	/*
			//------------------------------------------------------//
		                Setting up Auth layer
			//------------------------------------------------------//
	*/
	slog.Info("main : Started : Initializing authentication support")
	privatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePEM)
	if err != nil {
		panic(err)
	}

	pubKeyPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyPEM)
	if err != nil {
		panic(err)
	}

	k, err := auth.NewKeys(privateKey, publicKey)
	if err != nil {
		panic(err)
	}

	// setting up routes, and running the server
	r := gin.New()
	handlers.RegisterRoutes(r, conf, k)
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
