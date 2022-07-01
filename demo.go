package main

import (
	"MatchMove/common"
	"MatchMove/handler"
	pgquery "MatchMove/repository"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	host = getEnv("Postgres_Host", "127.0.0.1")
	port = string(getEnv("Postgres_Port", "5432"))
)

//Initialize intilaizes the server
func (dss *Demo) Initialize(user, password, dbname string) {

	handler := handler.NewHandler()

	//db connection
	var err error
	// host := os.Getenv("Postgres_Host")
	psqlconn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", psqlconn)

	common.CheckError(err)

	pgquery.Queryexec(db)
	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	common.CheckError(err)

	fmt.Println("router methods")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/v1/AdminUserLogin", handler.CheckUserInformation).Methods("POST")
	router.HandleFunc("/v1/ValidateToken", handler.ValidateToken).Methods("POST")
	router.HandleFunc("/v1/DisableToken", handler.DisableToken).Methods("PUT")

	fmt.Println("listening on 9294")
	log.Fatal(http.ListenAndServe(":9294", router))
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
