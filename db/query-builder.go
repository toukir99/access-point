package db

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var psql squirrel.StatementBuilderType

func InitQueryBuilder(writeDb *sqlx.DB) {
	psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	//createTables(writeDb)
}

// func createTables(writeDb *sqlx.DB) {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
//     id SERIAL PRIMARY KEY,
//     username VARCHAR(255) NOT NULL,
//     email VARCHAR(255) UNIQUE NOT NULL,
//     password VARCHAR(255) NOT NULL,
//     is_active BOOLEAN DEFAULT false,
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
// 	);`

// 	_, err := writeDb.Exec(query)
// 	if err != nil {
// 		log.Printf("Failed to create table 'users': %v", err)
// 	}
// }

func GetQueryBuilder() squirrel.StatementBuilderType {
	return psql
}
