package models

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" //get SQL driver
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Sql dialect
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error
var e error

func init() {
	err = godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	charset := os.Getenv("charset")
	parseTime := os.Getenv("parse_time")
	timezone := os.Getenv("timezone")

	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&&parseTime=%s&&loc=%s", username, password, dbHost, dbPort, dbName, charset, parseTime, timezone)
	conn, err := gorm.Open(dbType, dbURI)

	db = conn
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	fmt.Println("Connected to database on " + dbHost)

}

// GetDB .
func GetDB() *gorm.DB {
	return db
}
