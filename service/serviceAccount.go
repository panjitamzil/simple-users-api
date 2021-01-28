package service

import (
	"database/sql"
	"log"
	"users/domain/account/entity"
)

func (s *svc) InsertUser(acc entity.Account) error {
	err := s.account.Insert(acc)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *svc) UpdateUser(uid string, acc entity.Account) error {
	_, err := s.account.Get(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		log.Println(err)
		return err
	}

	err = s.account.Update(uid, acc)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *svc) DeleteUser(uid string) error {
	_, err := s.account.Get(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			return sql.ErrNoRows
		}
		log.Println(err)
		return err
	}

	err = s.account.Delete(uid)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *svc) GetAllUsers() (accounts []entity.AccountInfo, err error) {
	accounts, err = s.account.GetAll()
	if err != nil {
		log.Println(err)
		return accounts, err
	}

	return accounts, nil
}

func (s *svc) GetUser(uid string) (acc entity.AccountInfo, err error) {
	acc, err = s.account.Get(uid)
	if err != nil {
		log.Println(err)
		return acc, err
	}

	return acc, nil
}

func (s *svc) CheckUserID(uid string) bool {
	status := s.account.GetUserID(uid)
	if status != true {
		return false
	}

	return true
}
