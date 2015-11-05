package daos 

import (
    "database/sql"
)

var dbInfo string
var db *sql.DB

const (
    DB_USER     = "sc_admin"
    DB_PASSWORD = "sc_admin"
    DB_NAME     = "posts"
)


func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}