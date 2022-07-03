package server

import (
	"golibraryApi/database"
	"golibraryApi/handler"
)

func Start() {
	values := database.InitializeDbParameters()
	var PDB = new(database.PostgresDb)
	h := &handler.Handler{DB: PDB}

}
