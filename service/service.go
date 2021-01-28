package service

import (
	"users/domain/account/entity"
	repo "users/domain/account/repository"
)

// Service interface
type Service interface {
	InsertUser(entity.Account) error
	UpdateUser(uid string, acc entity.Account) error
	DeleteUser(uid string) error
	GetAllUsers() (accounts []entity.AccountInfo, err error)
	GetUser(uid string) (acc entity.AccountInfo, err error)
	CheckUserID(uid string) bool
}

type svc struct {
	account repo.Repository
}

// New ...
func New(
	account repo.Repository,
) Service {
	return &svc{
		account: account,
	}
}
