package database

import (
	"fmt"
	"golibraryApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDb struct {
	DB *gorm.DB
}

func (pdb *PostgresDb) SetupDb(host, user, password, dbName, port string) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, user, password, dbName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pdb.DB = db

	dberr := pdb.DB.AutoMigrate(&models.Book{})
	if dberr != nil {
		log.Fatal(dberr)
	}
	return nil
}

//GET ALL books FROM DB
func (pdb *PostgresDb) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := pdb.DB.Find(&books).Error; err != nil {
		log.Println("Could not find book", err)
	}

	return books, nil
}
