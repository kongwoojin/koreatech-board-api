package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"koreatech-board-api/cmd/utils"
	"log"
	"os"
)

var Pool = connect()

func connect() *pgxpool.Pool {
	if !utils.IsRunningInContainer() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")
	pgUser := os.Getenv("POSTGRES_USER")
	pgPasswd := os.Getenv("POSTGRES_PASSWORD")
	pgDBName := os.Getenv("POSTGRES_DB")

	// Construct connection string
	// Example: postgres://username:password@localhost:5432/database_name
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPasswd, pgHost, pgPort, pgDBName)

	ctx := context.Background()
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatalf("Unable to parse database config: %v\n", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// Verify connection
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	return pool
}
