package main

import (
	"github.com/joho/godotenv"
	"golibraryApi/router"
	"log"
)

func LoadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
func main() {
	library := router.SetupRouter()
	err := library.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
