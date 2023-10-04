package main

import (
	"database/sql"
	"math"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func openDB(dbDriver, dbSource string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB(dbDriver, dbSource string) *sql.DB {
	var connectionCounts int64
	var backOff = 1 * time.Second
	for {
		connection, err := openDB(dbDriver, dbSource)
		if err != nil {
			log.Info().Msg("Postgres not yet ready ...")
			connectionCounts++
		} else {
			log.Info().Msg("Connected to Postgres!")
			return connection
		}

		if connectionCounts > 10 {
			log.Fatal().Err(err)
			return nil
		}

		backOff = time.Duration(math.Pow(float64(connectionCounts), 2)) * time.Second
		log.Info().Msg("backing off...")
		time.Sleep(backOff)
		continue
	}
}
