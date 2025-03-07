package main

import (
	"fmt"
	"go-players/internal/api"
	"go-players/internal/database"
	"log"
	"net/http"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewServer(db)

	http.HandleFunc("/players", server.GetPlayers)

	fmt.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
