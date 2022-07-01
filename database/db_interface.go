package database

import (
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

//DbIface exposes db interconnection methods
type DbIface interface {
	DbExecuteScalar(query string, args ...interface{}) (pgx.Rows, error)
	TxBegin() (pgx.Tx, error)
	TxQuery(tx pgx.Tx, query string) (pgx.Rows, error)
	TxComplete(tx pgx.Tx) error
	TxCreateTempTable(tx pgx.Tx, tblName string, sql string) (*pgconn.StatementDescription, error)
	DbClose()
}
