package database

import (
	"github.com/joho/godotenv"
	"golibraryApi/models"
	"log"
	"os"
)

type DB interface {
	GetAllBooks() ([]models.Book, error)
	BookById(id string) (*models.Book, error)
	AddNewBook(book models.Book) error
	Checkout(id, copies string) (*models.Book, error)
	Checkin(id, copies string) (*models.Book, error)
	Deletebook(id string) error
}

type DbParameters struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func InitializeDbParameters() DbParameters {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	return DbParameters{
		Host:     host,
		User:     user,
		Password: password,
		DbName:   dbName,
		Port:     port,
	}
}
