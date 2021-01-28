package account

import "github.com/jmoiron/sqlx"

// MySQL struct
type MySQL struct {
	db *sqlx.DB
}

// New ...
func New(ds *sqlx.DB) *MySQL {
	return &MySQL{
		db: ds,
	}
}
