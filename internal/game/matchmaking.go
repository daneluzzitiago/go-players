package game

import (
	"go-players/internal/domain"
	"go-players/pkg/utils"
)

func matchmaking(players []domain.Player) (domain.Player, domain.Player) {
	var player1, player2 domain.Player
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			if utils.Abs(players[i].Level-players[j].Level) <= 5 {
				return players[i], players[j]
			}
		}
	}
	return player1, player2
}
