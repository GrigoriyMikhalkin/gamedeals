package database

import (
	"os"
	"fmt"
	"log"
	// "path/filepath"
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	// Global connection to DB via pop
	DB *gorm.DB
)


type configuration struct {
	DBUser 			string `json:"db_user"`
	DBPassword 	string `json:"db_password"`
	DBName 			string `json:"db_name"`
}


func InitDB() (*gorm.DB, error){
	var err error
	var file *os.File
	conf := configuration{}

	// Load DB configuration
	file, err = os.Open("dbconfig.json")
	if err != nil { log.Fatal(err) }

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&conf)
	if err != nil { log.Fatal(err) }

	// Connect to DB
	DB, err = gorm.Open(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=disable",
			conf.DBUser,
			conf.DBPassword,
			conf.DBName,
		),
	)
	if err != nil { log.Fatal(err) }

	return DB, err
}
