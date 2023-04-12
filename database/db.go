package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db_go_jwt_hacktiv8"
	db       *gorm.DB
	err      error
)

func InitDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.Debug().AutoMigrate()
}

func GetDB() *gorm.DB {
	return db
}
