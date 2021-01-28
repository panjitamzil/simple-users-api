package repository

import "users/domain/account/entity"

// Repository struct
type Repository interface {
	// reader
	GetAll() (accounts []entity.AccountInfo, err error)
	Get(uid string) (acc entity.AccountInfo, err error)
	GetUserID(uid string) bool

	// writer
	Insert(acc entity.Account) error
	Update(uid string, acc entity.Account) error
	Delete(uid string) error
}
