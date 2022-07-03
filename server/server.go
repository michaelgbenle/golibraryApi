package server

import (
	"golibraryApi/database"
	"golibraryApi/handler"
	"log"
)

func Start() error {
	values := database.InitializeDbParameters()
	var PDB = new(database.PostgresDb)
	h := &handler.Handler{DB: PDB}

	err := PDB.SetupDb(values.Host, values.User, values.Password, values.DbName, values.Port)
	if err != nil {
		log.Println("Error trying to Init", err)
		return err
	}
	return nil
}
