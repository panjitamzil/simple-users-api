package account

import (
	"database/sql"
	"log"
	"users/domain/account/entity"
)

// GetAll to get all accounts
func (m *MySQL) GetAll() (accounts []entity.AccountInfo, err error) {
	var acc []entity.Account

	err = m.db.Select(&acc, `
		SELECT id, user_id, email, address, password, created_at, updated_at, deleted_at
		FROM accounts
		WHERE status = 1
	`)

	if err != nil {
		log.Println(err)
		return accounts, err
	}

	for _, a := range acc {
		accounts = append(accounts, entity.AccountInfo{
			ID:        a.ID,
			UserID:    a.UserID,
			Email:     a.Email,
			Address:   a.Address,
			Password:  a.Password,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
		})
	}

	if len(accounts) == 0 {
		return accounts, sql.ErrNoRows
	}

	return accounts, nil
}

// Get to get account by id
func (m *MySQL) Get(id string) (acc entity.AccountInfo, err error) {
	var a entity.Account

	err = m.db.Get(&a, `
		SELECT id, user_id, email, address, password, created_at, updated_at, deleted_at
		FROM accounts
		WHERE user_id = ? AND status = 1
	`, id)

	if err != nil {
		log.Println(err)
		return acc, err
	}

	acc = entity.AccountInfo{
		ID:        a.ID,
		UserID:    a.UserID,
		Email:     a.Email,
		Address:   a.Address,
		Password:  a.Password,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		DeletedAt: a.DeletedAt,
	}

	return acc, nil
}

// GetUserID to get user id from db
func (m *MySQL) GetUserID(uid string) bool {
	var a string

	err := m.db.Get(&a, `
		SELECT user_id
		FROM accounts
		WHERE user_id = ? AND status = 1
	`, uid)

	if err != nil {
		if err == sql.ErrNoRows {
			return true
		}
		log.Println(err)
		return false
	}

	return false
}
