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

type Tournament struct {
	name    string
	id      int
	players []Player
	winner  Player
}

var players []Player

func (t *Tournament) addPlayer(player Player) {
	t.players = append(t.players, player)
}

func (t *Tournament) playTournament() Player {
	for len(t.players) > 1 {
		player1, player2 := t.players[0], t.players[1]
		winner := playMatch(player1, player2)
		t.players = append([]Player{winner}, t.players[2:]...)
	}
	t.winner = t.players[0]
	return t.winner
}

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
	tournament := Tournament{name: "Amazing Tournament", id: 1}
	addPlayer("Player1", 1, 10, 0, 0)
	addPlayer("Player2", 2, 17, 0, 0)
	addPlayer("Player3", 3, 14, 0, 0)

	tournament.addPlayer(players[0])
	tournament.addPlayer(players[1])
	tournament.addPlayer(players[2])

	fmt.Println("Player list:")
	listPlayers()

	fmt.Println("Tournament players:")
	for _, player := range tournament.players {
		fmt.Printf("(%d) - %s, Level: %d, Wins: %d, Losses: %d\n",
			player.id, player.name, player.level, player.wins, player.losses)
	}

	winner := tournament.playTournament()

	fmt.Printf("The winner of the tournament %s is %s!\n", tournament.name, winner.name)
}
