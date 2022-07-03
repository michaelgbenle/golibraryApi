package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDb struct {
	DB *gorm.DB
}

func SetupDb() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
