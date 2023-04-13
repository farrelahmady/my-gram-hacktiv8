package database

import (
	"fmt"
	"log"

	"github.com/farrelahmady/my-gram-hacktiv8/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db_mygram_hacktiv8"
	db       *gorm.DB
	err      error
)

func InitDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.Debug().AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.SocialMedia{},
	)
}

func GetDB() *gorm.DB {
	return db
}
