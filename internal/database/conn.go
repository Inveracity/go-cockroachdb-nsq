package crud

import (
	"log"
	"time"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // nolint:gci // sql driver
	_ "github.com/jackc/pgx/v4"                                // nolint:gci // sql driver
	"github.com/jmoiron/sqlx"
)

// CreateDatabase attempts to create a database, if database already exists the error
// will be ignored
func CreateDatabase(dbName string) (*sqlx.DB, error) {
	db, err := Database()
	if err != nil {
		return db, err
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)

	return nil, err
}

func CreateTable(dbName string) (*sqlx.DB, error) {
	db, err := Database()
	if err != nil {
		return db, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id UUID NOT NULL,
		os STRING NOT NULL,
		version STRING NOT NULL
	);`)

	return nil, err
}

// Database creates a new sqlx.DB object from the environment variables
// to manage and maintain database connections
func Database() (*sqlx.DB, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Minute

	var (
		dbconn *sqlx.DB
		err    error
	)

	operation := func() error {
		log.Print("connecting to the database")

		dbconn, err = sqlx.Connect("postgres", "host=localhost port=23257 user=root dbname=work sslmode=disable")
		if err != nil {
			log.Printf("failed to connect to the database: %v. retrying in %s", err, bo.NextBackOff())
		}

		return err
	}

	err = backoff.Retry(operation, bo)
	if err != nil {
		return nil, err
	}

	return dbconn, nil
}
