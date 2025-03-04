package main

import "fmt"

type Player struct {
	Name   string
	ID     int
	Level  int
	Wins   int
	Losses int
}

var Players []Player

func addNewPlayer(name string, id int) {
	Players = append(Players, Player{name, id, 1, 0, 0})
}

func listPlayers() {
	for _, player := range Players {
		fmt.Printf("(%d) - %s, Level: %d, Wins: %d, Losses: %d\n",
			player.ID, player.Name, player.Level, player.Wins, player.Losses)
	}
}

func main() {
	addNewPlayer("Player 1", 1)
	addNewPlayer("Player 2", 2)

	fmt.Println("Player list:")
	listPlayers()
}
