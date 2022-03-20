package crud

import (
	"context"

	"github.com/inveracity/go-cockroachdb-nsq/internal/task"
	_ "github.com/jackc/pgx/v4" // nolint:gci // sql driver
	"github.com/jmoiron/sqlx"
)

// DB is the database layer for work database
type DB struct {
	db *sqlx.DB
}

// NewDB creates a new instance of DB
func NewDB(db *sqlx.DB) *DB {
	return &DB{db}
}

// Insert inserts a new row in the tasks table
func (db *DB) Insert(ctx context.Context, task task.Task) error {
	sqlStmt := `
		INSERT INTO tasks (
			id,
			os,
			version
		) VALUES (
			:id,
			:os,
			:version
		)`

	args := map[string]interface{}{
		"id":      task.ID.String(),
		"os":      task.Os,
		"version": task.Version,
	}

	_, err := sqlx.NamedExecContext(ctx, db.db, sqlStmt, args)

	return err
}

// Read takes a task id and returns a task
func (db *DB) Read(ctx context.Context) (task.Task, error) {
	taskFromDB := task.Task{}

	sqlStmt := `SELECT * FROM tasks LIMIT 1`

	args := map[string]interface{}{}

	prepStmt, err := db.db.PrepareNamedContext(ctx, sqlStmt)
	if err != nil {
		return task.Task{}, err
	}

	err = prepStmt.GetContext(ctx, &taskFromDB, args)
	if err != nil {
		return task.Task{}, err
	}

	return taskFromDB, nil
}
