package handler

import (
	"net/http"
)

//Ihandler provides the abstraction
type Ihandler interface {
	CheckUserInformation(w http.ResponseWriter, r *http.Request)
	ValidateToken(w http.ResponseWriter, r *http.Request)
	DisableToken(w http.ResponseWriter, r *http.Request)
}