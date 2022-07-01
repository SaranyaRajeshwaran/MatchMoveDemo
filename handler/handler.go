package handler

import (
	model "MatchMove/model"
	"MatchMove/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var secretkey = []byte("secret_key")

//Handler provies the Boz lgoic before calling the service
type Handler struct {
	DatabaseService service.IPSQLService
}

//NewHandler is teh ctor for the hadnler
func NewHandler() Ihandler {
	return &Handler{
		DatabaseService: service.NewDatabaseServicesInstance(),
	}
}

//CheckUserInformation-Checks user Information
func (handler *Handler) CheckUserInformation(w http.ResponseWriter, r *http.Request) {
	// Unmarshall json values
	log.Println("CheckUserInformation")
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	var login model.Login
	err = json.Unmarshal(request, &login)

	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err)
		return
	}
	isExists, err := handler.DatabaseService.CheckUserInformation(&login)
	if err != nil {

		log.Println(err)
		responseController(w, http.StatusOK, isExists)
		return
	}
	log.Println("isExists")
	log.Println(isExists)

	responseController(w, http.StatusOK, isExists)
}

//ValidateToken- validates token for admin user
func (handler *Handler) ValidateToken(w http.ResponseWriter, r *http.Request) {
	// Unmarshall json values
	log.Println("VDiableToken")
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	var token *model.ValidateToken
	err = json.Unmarshal(request, &token)

	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err)
		return
	}
	isExists, err := handler.DatabaseService.ValidateToken(token.Token)
	if err != nil {

		log.Println(err)
		responseController(w, http.StatusOK, isExists)
		return
	}
	log.Println("isExists")
	log.Println(isExists)
	responseController(w, http.StatusOK, isExists)

}

//DisableToken- Disables token for admin user
func (handler *Handler) DisableToken(w http.ResponseWriter, r *http.Request) {
	// Unmarshall json values
	log.Println("ValidateToken")
	request, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}
	var token *model.ValidateToken
	err = json.Unmarshal(request, &token)

	if err != nil {
		log.Println(err)
		responseController(w, http.StatusInternalServerError, err)
		return
	}
	isExists, err := handler.DatabaseService.DisableToken(token.UserName, token.Token)
	if err != nil {

		log.Println(err)
		responseController(w, http.StatusOK, isExists)
		return
	}
	log.Println("isExists")
	log.Println(isExists)
	responseController(w, http.StatusOK, isExists)
}

//Define response controller
func responseController(w http.ResponseWriter, code int, payload interface{}) {
	enableCors(&w)
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
