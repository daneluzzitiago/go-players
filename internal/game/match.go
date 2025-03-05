package game

import (
	"fmt"
	"go-players/internal/domain"
	"math/rand"
)

func PlayMatch(player1 domain.Player, player2 domain.Player) domain.Player {
	winner := player1
	if rand.Intn(2) == 0 {
		winner = player2
	}
	fmt.Printf("The match between %s and %s is over! The winner is %s!\n", player1.Name, player2.Name, winner.Name)
	return winner
}
