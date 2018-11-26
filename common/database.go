package common

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"

	// postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

// Database struct {
type Database struct {
	*gorm.DB
}

func getDBConfig() (string, error) {
	host, ok := os.LookupEnv("HOST")
	if !ok {
		return "", errors.New("empty hostname")
	}

	user, ok := os.LookupEnv("USER")
	if !ok {
		return "", errors.New("empty user")
	}

	password, ok := os.LookupEnv("PASSWORD")
	if !ok {
		return "", errors.New("empty password")
	}

	hash := "$2a$14$o71yt2NdDJFD/HBj2HHsjusYq7ndOwA5w9PAF09dkno.Tlz2i/tMW"

	match := CheckPasswordHash(password, hash)
	if !match {
		return "", errors.New("password incorrect.")
	}

	dbname, ok := os.LookupEnv("DBNAME")
	if !ok {
		return "", errors.New("empty dbname")
	}

	portN, ok := os.LookupEnv("PORT")
	if !ok {
		return "", errors.New("empty port")
	}
	port, err := strconv.Atoi(portN)
	if err != nil {
		return "", errors.New("Invalid port number")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	return psqlInfo, nil
}

// DB *gorm.DB
var DB *gorm.DB

// Init Opening a database and save the reference to `Database` struct.
func Init() *gorm.DB {
	psqlInfo, perr := getDBConfig()
	if perr != nil {
		log.Println("DBConfig error:", perr)
		os.Exit(1)
	}

	log.Println("DB Connecting...")
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("DB Open err: ", err)
	}
	db.DB().SetMaxIdleConns(10)
	// db.LogMode(true)
	DB = db
	return DB
}

// TestDBInit This function will create a temporarily database for running testing cases
func TestDBInit() *gorm.DB {
	psqlInfo, perr := getDBConfig()
	if perr != nil {
		return nil
	}

	testDb, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("db err: ", err)
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
