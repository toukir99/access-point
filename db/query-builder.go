package db

import (
	"log"

	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var psql squirrel.StatementBuilderType

// InitQueryBuilder initializes the Squirrel statement builder and creates tables
func InitQueryBuilder(writeDb *sqlx.DB) {
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	// Create necessary tables
	createTables(writeDb)
}

// createTables creates the necessary tables in the database using raw SQL
func createTables(writeDb *sqlx.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := writeDb.Exec(query)
	if err != nil {
		log.Printf("Failed to create table 'users': %v", err)
	}
}

// GetQueryBuilder returns the initialized Squirrel statement builder
func GetQueryBuilder() squirrel.StatementBuilderType {
	return psql
}
