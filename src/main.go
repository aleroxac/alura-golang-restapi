package main

import (
	"fmt"

	"github.com/aleroxac/alura-golang/database"
	"github.com/aleroxac/alura-golang/routes"
)

func main() {
	database.InitDatabaseConnection()
	fmt.Println("Listening on :8000")
	routes.HandleRequest()
}
