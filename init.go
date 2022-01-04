package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/mig"
)

var user = os.Getenv("SQL_USER")
var pass = os.Getenv("SQL_PASS")
var dbName = os.Getenv("SQL_DB")
var sqlHost = os.Getenv("SQL_HOST")
var sqlPort = os.Getenv("SQL_PORT")
var port = os.Getenv("PORT")
var pool *sql.DB
var migrationPath = os.Getenv("SQL_MIGS")

func init() {
	var err error
	sqlConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, sqlHost, sqlPort, dbName)
	pool, err = sql.Open("mysql", sqlConnectionString)
	if err != nil {
		panic(err)
	}
	_, err = mig.Up("mysql", sqlConnectionString, migrationPath)
	if err != nil {
		panic(err)
	}
}
