package rest

import (
	"net/http"
	"users/service"

	"github.com/julienschmidt/httprouter"
)

// Rest struct
type Rest struct {
	service service.Service
}

// New ...
func New(
	service service.Service,
) *Rest {
	return &Rest{
		service: service,
	}
}

// Routing is list of endpoint
func (rest *Rest) Routing(router *httprouter.Router) {
	router.POST("/register", rest.handleRegister)
	router.GET("/profile/:uid", rest.handleGetProfile)
	router.PUT("/profile/update/:uid", rest.handleUpdateProfile)
	router.GET("/users", rest.handleGetUser)
	router.DELETE("/user/delete/:uid", rest.handleDeleteUser)

	router.POST("/login", rest.handleLogin)
	router.GET("/ping", handlePing)

}

func handlePing(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("PONG"))
}
