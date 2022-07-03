package database

import (
	"fmt"
	"golibraryApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strconv"
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
	if err := pdb.DB.Where("ID = ?", id).Find(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (pdb *PostgresDb) AddNewBook(book models.Book) error {
	if err := pdb.DB.Create(&book).Error; err != nil {
		return err
	}
	return nil
}

func (pdb *PostgresDb) Checkout(id, copies string) (*models.Book, error) {
	book := &models.Book{}
	intCopies, _ := strconv.Atoi(copies)
	newId, _ := strconv.Atoi(id)

	if err := pdb.DB.Model(book).Where("id", newId).Update("quantity", book.Quantity-intCopies).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (pdb *PostgresDb) Checkin(id, copies string) (*models.Book, error) {
	var book *models.Book
	intCopies, _ := strconv.Atoi(copies)
	newQuantity := book.Quantity + intCopies
	if err := pdb.DB.Model(book).Where(book.ID, id).Update(strconv.Itoa(book.Quantity), newQuantity).Error; err != nil {
		return nil, err
	}
	return book, nil
}
