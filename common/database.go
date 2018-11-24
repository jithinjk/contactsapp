package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Database struct {
type Database struct {
	*gorm.DB
}

const (
	host     = "dumbo.db.elephantsql.com"
	port     = 5432
	user     = "zokkzkcw"
	password = "UKkRB_MI6AB-pJJ6ZpULaBOdL7gNITw8"
	dbname   = "zokkzkcw"
)

// DB *gorm.DB
var DB *gorm.DB

// Init Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	// db.LogMode(true)
	DB = db
	return DB
}

// TestDBInit This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s", host, port, user, dbname, password)

	testDb, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	testDb.DB().SetMaxIdleConns(3)
	testDb.LogMode(true)
	DB = testDb
	return DB
}

// GetDB Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
