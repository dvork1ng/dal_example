package core

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DBConn *sql.DB

func InitDatabase(host string, port int, user string, pass string, dbName string) error {
	var err error
	DBConn, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, pass))
	if err != nil {
		DBConn = nil
		return err
	}
	return nil
}
