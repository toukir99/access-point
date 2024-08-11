package db

import (
	"log/slog"
	"os"
	"access-point/config"

	_ "github.com/lib/pq"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB() {
	conf := config.GetConfig()

	migrations := &migrate.FileMigrationSource{
		Dir: conf.MigrationSource,
	}

	_, err := migrate.Exec(writeDb.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("Successfully migrated database!")
}