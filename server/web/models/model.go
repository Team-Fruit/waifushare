package models

package models

import (
	"github.com/jmoiron/sqlx"
)

type (
	Model struct {
		db *sqlx.DB
	}

	Database interface {
	}
)

func NewModel(db *sqlx.DB) *Model {
	return &Model{
		db: db,
	}
}
