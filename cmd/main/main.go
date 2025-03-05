package main

import (
	"fmt"
	"go-players/internal/database"
	"log"
)

func main() {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	players, err := db.GetPlayers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nPlayer List:")
	fmt.Println("----------------------------------------")
	for _, p := range players {
		fmt.Printf("%-25s | Level: %2d | W/L: %2d/%2d\n",
			p.Name,
			p.Level,
			p.Wins,
			p.Losses,
		)
	}
	fmt.Println("----------------------------------------")
}
