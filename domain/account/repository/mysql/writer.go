package account

import (
	"database/sql"
	"log"
	"time"
	"users/domain/account/entity"
)

// Insert to insert data into database
func (m *MySQL) Insert(acc entity.Account) error {
	var (
		args []interface{}
		now  = time.Now()
	)

	args = append(args, acc.UserID, acc.Email, acc.Address, acc.Password, now)

	query := `INSERT INTO accounts (user_id, email, address, password, created_at, status) VALUES (?,?,?,?,?,1)`

	_, err := m.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// Update to update data
func (m *MySQL) Update(id string, acc entity.Account) error {
	var (
		args []interface{}
		now  = time.Now()
	)

	args = append(args, acc.UserID, acc.Email, acc.Address, acc.Password, now, id)

	query := `UPDATE accounts
				SET 
			user_id = ?,
			email = ?,
			address = ?,
			password = ?,
			updated_at = ?
				WHERE user_id = ? AND status = 1`

	r, err := m.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	num, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// Delete to delete data
func (m *MySQL) Delete(id string) error {
	var (
		args []interface{}
		now  = time.Now()
	)

	args = append(args, now, id)

	query := `UPDATE accounts
				SET status = 0, deleted_at = ?
				WHERE user_id = ? AND status = 1`

	r, err := m.db.Exec(query, args...)
	if err != nil {
		log.Println(err)
		return err
	}

	num, err := r.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}

	if num == 0 {
		return sql.ErrNoRows
	}

	return nil
}
