package error

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Error struct
type Error struct {
	Status int    `json:"status"`
	Title  string `json:"title"`
}

var (
	defaultErrorResp []byte
	// ErrInvalidRequest ...
	ErrInvalidRequest = errors.New("You have invalid request, please check your request again")
	// ErrInternalServerError ...
	ErrInternalServerError = errors.New("Internal Server Error! We will fix it")
	// ErrEntityNotFound ...
	ErrEntityNotFound = errors.New("Sorry, data not found")
	// ErrBadRequest ...
	ErrBadRequest = errors.New("Bad request! Username or Password unmatch")
	// ErrExistUserID ...
	ErrExistUserID = errors.New("User id already used! Please use another user id")
)

// WithError to return error on json format
func WithError(w http.ResponseWriter, httpStatusCode int, errs error) {
	resp := map[string]interface{}{
		"errors": Error{
			Status: httpStatusCode,
			Title:  errs.Error(),
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		jsonResp = defaultErrorResp
		httpStatusCode = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	w.Write([]byte(jsonResp))
}

// WithData to return data into json format
func WithData(w http.ResponseWriter, data interface{}) {
	status := http.StatusOK
	resp := map[string]interface{}{
		"data": data,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		jsonResp = defaultErrorResp
		status = http.StatusInternalServerError
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(jsonResp))
}
