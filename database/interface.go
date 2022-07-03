package database

import "os"

type DB interface {
}

type DbParameters struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}

func InitializeDbParameters() DbParameters {
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