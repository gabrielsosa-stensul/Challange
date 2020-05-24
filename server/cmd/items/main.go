package main

import (
	"log"

	"github.com/MarianoArias/Challange/server/internal/app/items/api"
)

// @title Api Go
// @version 1.0
// @description Challange

// @contact.name Mariano Arias
// @contact.url www.github.com/MarianoArias
// @contact.email mariano.arias.1987@gmail.com

// @host localhost:8080
// @BasePath /
func main() {
	router := api.SetupRouter()
	log.Printf("\033[97;42m%s\033[0m\n", "Ready to go :)")
	router.Run(":8080")
}
