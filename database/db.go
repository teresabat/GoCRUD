// database/database.go
package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
    var err error
    DB, err = sql.Open("mysql", "root:0000@tcp(127.0.0.1:3306)/crud_golang")
    if err != nil {
        panic(err.Error())
    }
}
