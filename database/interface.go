package database

type DB interface {
}

type DbParameters struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     string
}
