package domain

import "fmt"

type Player struct {
	Name   string
	ID     int
	Level  int
	Wins   int
	Losses int
}

var Players []Player

func AddNewPlayer(name string, id int) {
	Players = append(Players, Player{name, id, 1, 0, 0})
}

func AddPlayer(name string, id int, level int, wins int, losses int) {
	Players = append(Players, Player{name, id, level, wins, losses})
}

func ListPlayers() {
	for _, player := range Players {
		fmt.Printf("(%d) - %s, Level: %d, Wins: %d, Losses: %d\n",
			player.ID, player.Name, player.Level, player.Wins, player.Losses)
	}
}
