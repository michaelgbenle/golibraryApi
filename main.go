package main

import (
	"github.com/joho/godotenv"
	"golibraryApi/router"
	"log"
	"os"
)

func main() {
	everr := godotenv.Load(".env")
	if everr != nil {
		log.Fatal(everr)
	}
	port := os.Getenv("PORT")

	library := router.SetupRouter()
	err := library.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
