package database

import (
	"golibraryApi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type PostgresDb struct {
	DB *gorm.DB
}

//Export an instance of database
func NewPostgresDb() *PostgresDb {
	return &PostgresDb{}
}
func (pdb *PostgresDb) SetupDb() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	pdb.DB = db

	err = pdb.DB.AutoMigrate(models.Book{})
}
