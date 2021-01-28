package account

import "time"

type (
	// Account struct
	Account struct {
		UserID    string    `db:"user_id"`
		Email     string    `db:"email"`
		Address   string    `db:"address"`
		Password  string    `db:"password"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		DeletedAt time.Time `db:"deleted_at"`
	}
)
