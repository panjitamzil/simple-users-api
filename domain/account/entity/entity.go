package entity

import (
	"gopkg.in/guregu/null.v3"
)

type (
	// Account struct
	Account struct {
		ID        int64     `db:"id"`
		UserID    string    `db:"user_id"`
		Email     string    `db:"email"`
		Address   string    `db:"address"`
		Password  string    `db:"password"`
		CreatedAt null.Time `db:"created_at"`
		UpdatedAt null.Time `db:"updated_at"`
		DeletedAt null.Time `db:"deleted_at"`
	}

	// AccountInfo struct
	AccountInfo struct {
		ID        int64     `json:"id"`
		UserID    string    `json:"user_id"`
		Email     string    `json:"email"`
		Address   string    `json:"address"`
		Password  string    `json:"password"`
		CreatedAt null.Time `json:"created_at,omitempty"`
		UpdatedAt null.Time `json:"updated_at,omitempty"`
		DeletedAt null.Time `json:"deleted_at,omitempty"`
	}

	// Login struct
	Login struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)
