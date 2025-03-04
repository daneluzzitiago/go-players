package main

import (
	"fmt"
	"math/rand"
)

type Player struct {
	name   string
	id     int
	level  int
	wins   int
	losses int
}

var players []Player

func addNewPlayer(name string, id int) {
	players = append(players, Player{name, id, 1, 0, 0})
}

func addPlayer(name string, id int, level int, wins int, losses int) {
	players = append(players, Player{name, id, level, wins, losses})
}

func listPlayers() {
	for _, player := range players {
		fmt.Printf("(%d) - %s, Level: %d, Wins: %d, Losses: %d\n",
			player.id, player.name, player.level, player.wins, player.losses)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func matchmaking() (Player, Player) {
	var players1, player2 Player
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if abs(players[i].level-players[j].level) <= 5 {
				players1 = players[i]
				player2 = players[j]
				return players1, player2
			}
		}
	}
	return players1, player2
}

func playMatch(player1 Player, player2 Player) Player {
	winner := player1
	if rand.Intn(2) == 0 {
		winner = player2
	}
	fmt.Printf("The match between %s and %s is over! The winner is %s!\n", player1.name, player2.name, winner.name)
	return winner
}

func main() {
	addPlayer("Player1", 1, 10, 0, 0)
	addPlayer("Player2", 2, 17, 0, 0)
	addPlayer("Player3", 3, 14, 0, 0)

	fmt.Println("Player list:")
	listPlayers()

	player1, player2 := matchmaking()
	winner := playMatch(player1, player2)
	fmt.Println("Winner:", winner.name)
}
