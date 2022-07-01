package pgquery

import (
	"MatchMove/common"
	"bytes"
	"database/sql"
	_ "embed"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/qustavo/dotsql"
)

//go:embed init.psql
var initPSQL []byte

func Queryexec(db *sql.DB) {

	// Loads queries from file
	dot, err := dotsql.Load(bytes.NewReader(initPSQL))
	fmt.Println("load file")
	common.CheckError(err)

	// Run queries
	_, err = dot.Exec(db, "Demo-UserInformation-table")
	fmt.Println("Created Demo.USerInformation-table")
	common.CheckError(err)

	// Run queries
	_, err = dot.Exec(db, "Demo-Token-table")
	fmt.Println("Created Demo.Token-table")
	common.CheckError(err)

	_, err = dot.Exec(db, "create-DEMO-UserInformation")
	common.CheckError(err)
	fmt.Println("Created Demo.USerInformation-table with default values")

	fmt.Println("DB Connected!")
}
