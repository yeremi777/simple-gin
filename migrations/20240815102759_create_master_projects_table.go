package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateMasterProjectsTable, downCreateMasterProjectsTable)
}

func upCreateMasterProjectsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE master_projects (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			url VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

func downCreateMasterProjectsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `DROP TABLE IF EXISTS master_projects`)
	return err
}
