package server

import "golibraryApi/database"

func Start() {
	values := database.InitializeDbParameters()
	var PDB = new(database.PostgresDb)
}
