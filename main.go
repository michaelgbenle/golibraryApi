package main

import (
	"golibraryApi/router"
	"log"
)

func main() {
	library := router.SetupRouter()
	err := library.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
