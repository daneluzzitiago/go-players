package main

import (
	"fmt"
	"go-players/internal/domain"
	"go-players/internal/game"
)

func main() {
	tournament := domain.Tournament{Name: "Amazing Tournament", ID: 1}
	domain.AddPlayer("Player1", 1, 10, 0, 0)
	domain.AddPlayer("Player2", 2, 17, 0, 0)
	domain.AddPlayer("Player3", 3, 14, 0, 0)

	for _, player := range domain.Players {
		tournament.AddPlayer(player)
	}

	fmt.Println("Playexr list:")
	domain.ListPlayers()

	fmt.Println("Tournament players:")
	for _, player := range tournament.Players {
		fmt.Printf("(%d) - %s, Level: %d, Wins: %d, Losses: %d\n",
			player.ID, player.Name, player.Level, player.Wins, player.Losses)
	}

	winner := tournament.PlayTournament(game.PlayMatch)

	fmt.Printf("The winner of the tournament %s is %s!\n", tournament.Name, winner.Name)
}
