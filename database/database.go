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

// GetAllBooks gets ALL books FROM DB
func (pdb *PostgresDb) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := pdb.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (pdb *PostgresDb) BookById(id string) (*models.Book, error) {
	var book models.Book
	if err := pdb.DB.Where("ID = ?", id).First(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (pdb *PostgresDb) AddNewBook(book models.Book) error {
	if err := pdb.DB.Create(book).Error; err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDb) Checkout(id, copies string) (models.Book, error) {
	var book models.Book
	if err := pdb.DB.Model(book).Where("id = ?", id).Update()
}
