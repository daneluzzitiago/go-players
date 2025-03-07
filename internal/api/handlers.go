package api

import (
	"encoding/json"
	"go-players/internal/database"
	"log"
	"net/http"
)

type Server struct {
	db *database.Database
}

func NewServer(db *database.Database) *Server {
	return &Server{db: db}
}

func (s *Server) GetPlayers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	players, err := s.db.GetPlayers()
	if err != nil {
		log.Printf("Error fetching players: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}
