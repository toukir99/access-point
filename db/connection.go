package db

import (
	"access-point/config"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	readDb  *sqlx.DB
	WriteDb *sqlx.DB
)

func connect(dbConfig config.DBConfig) *sqlx.DB {
	dbSource := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	if !dbConfig.EnableSSLMode {
		dbSource += " sslmode=disable"
	}

	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dbCon.SetConnMaxIdleTime(
		time.Duration(dbConfig.MaxIdleTimeInMinute * int(time.Minute)),
	)

	return dbCon
}

func ConnectDB() {
	conf := config.GetConfig()
	readDb = connect(conf.DB.Read) //readDb.Exec("serfojj")
	if readDb != nil {
		slog.Info("Connected to read database!")
	}

	WriteDb = connect(conf.DB.Write)
	if WriteDb != nil {
		slog.Info("Connected to write database!")
	}
}

func CloseDB() {
	if err := readDb.Close(); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Disconnected from read database!")

	if err := WriteDb.Close(); err != nil {
		slog.Error(err.Error())
	}
	slog.Info("Disconnected from write database!")
}

func GetWriteDB() *sqlx.DB {
	return WriteDb
}

func GetReadDB() *sqlx.DB {
	return readDb
}
