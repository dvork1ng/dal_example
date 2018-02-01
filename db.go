package main

import (
	"fmt"

	"github.com/AplaProject/go-apla/packages/conf"
	"github.com/jinzhu/gorm"
)

var DBConn *gorm.DB

func GormInit(host string, port int, user string, pass string, dbName string) error {
	var err error
	DBConn, err = gorm.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbName, pass))
	if err != nil {
		DBConn = nil
		return err
	}
	if *conf.LogSQL {
		DBConn.LogMode(true)
	}
	return nil
}
