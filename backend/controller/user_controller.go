package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserController interface {
	Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ForgotPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	ResetPassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
