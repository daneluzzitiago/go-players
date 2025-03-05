package domain

type Tournament struct {
	Name    string
	ID      int
	Players []Player
	Winner  Player
}

func (t *Tournament) AddPlayer(player Player) {
	t.Players = append(t.Players, player)
}

func (t *Tournament) PlayTournament(playMatch func(Player, Player) Player) Player {
	for len(t.Players) > 1 {
		player1, player2 := t.Players[0], t.Players[1]
		winner := playMatch(player1, player2)
		t.Players = append([]Player{winner}, t.Players[2:]...)
	}
	t.Winner = t.Players[0]
	return t.Winner
}
