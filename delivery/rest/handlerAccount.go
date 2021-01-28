package rest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"users/domain/account/entity"
	errors "users/lib/error"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

func (rest *Rest) handleRegister(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := &entity.AccountInfo{}
	ParseBody(r, data)

	// check is it user_id exist or not
	status := rest.service.CheckUserID(data.UserID)
	if status != true {
		errors.WithError(w, http.StatusInternalServerError, errors.ErrExistUserID)
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	d := entity.Account{
		UserID:   data.UserID,
		Email:    data.Email,
		Address:  data.Address,
		Password: string(pass),
	}

	err = rest.service.InsertUser(d)
	if err != nil {
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	errors.WithData(w, d)
}

func (rest *Rest) handleUpdateProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uid := params.ByName("uid")

	data := &entity.AccountInfo{}
	ParseBody(r, data)

	pass, err := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	if err != nil {
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	d := entity.Account{
		UserID:   data.UserID,
		Email:    data.Email,
		Address:  data.Address,
		Password: string(pass),
	}

	err = rest.service.UpdateUser(uid, d)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	errors.WithData(w, d)
}

func (rest *Rest) handleGetProfile(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uid := params.ByName("uid")

	userInfo, err := rest.service.GetUser(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	errors.WithData(w, userInfo)
}

func (rest *Rest) handleGetUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	usersInfo, err := rest.service.GetAllUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	errors.WithData(w, usersInfo)
	return
}

func (rest *Rest) handleDeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	uid := params.ByName("uid")

	// check user has been created or not
	_, err := rest.service.GetUser(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	err = rest.service.DeleteUser(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	errors.WithData(w, "OK")
}

func (rest *Rest) handleLogin(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := &entity.Login{}
	ParseBody(r, data)

	// get users data
	UserInfo, err := rest.service.GetUser(data.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			errors.WithError(w, http.StatusNotFound, errors.ErrEntityNotFound)
			return
		}
		errors.WithError(w, http.StatusInternalServerError, errors.ErrInternalServerError)
		return
	}

	// check is it user's password and password from request match or not
	status := CheckPasswordHash(data.Password, UserInfo.Password)
	if status != true {
		errors.WithError(w, http.StatusBadRequest, errors.ErrBadRequest)
		return
	}

	success := fmt.Sprintf("Welcome %s", UserInfo.UserID)
	errors.WithData(w, success)
}

// ParseBody to parse body request into struct
func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// CheckPasswordHash to compare password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
