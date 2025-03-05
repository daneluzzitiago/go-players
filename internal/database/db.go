package database

import (
	"database/sql"
	"fmt"
	"go-players/internal/domain"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("sqlite3", "./players.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createTables(db); err != nil {
		return nil, err
	}

	database := &Database{db: db}

	if err := database.SeedDatabase(); err != nil {
		return nil, fmt.Errorf("error seeding database: %v", err)
	}

	return database, nil
}

func createTables(db *sql.DB) error {
	playerTable := `CREATE TABLE IF NOT EXISTS players (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        level INTEGER NOT NULL,
        wins INTEGER NOT NULL,
        losses INTEGER NOT NULL,
        deleted BOOLEAN DEFAULT FALSE
    );`

	tournamentTable := `CREATE TABLE IF NOT EXISTS tournaments (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        winner_id INTEGER,
        deleted BOOLEAN DEFAULT FALSE,
        FOREIGN KEY(winner_id) REFERENCES players(id)
    );`

	tournamentPlayersTable := `CREATE TABLE IF NOT EXISTS tournament_players (
        tournament_id INTEGER,
        player_id INTEGER,
        deleted BOOLEAN DEFAULT FALSE,
        PRIMARY KEY (tournament_id, player_id),
        FOREIGN KEY (tournament_id) REFERENCES tournaments(id),
        FOREIGN KEY (player_id) REFERENCES players(id)
    );`

	if _, err := db.Exec(playerTable); err != nil {
		return err
	}

	if _, err := db.Exec(tournamentTable); err != nil {
		return err
	}

	if _, err := db.Exec(tournamentPlayersTable); err != nil {
		return err
	}

	return nil
}

func (d *Database) SavePlayer(player domain.Player) error {
	stmt := `INSERT INTO players (id, name, level, wins, losses) 
             VALUES (?, ?, ?, ?, ?)`

	_, err := d.db.Exec(stmt, player.ID, player.Name, player.Level,
		player.Wins, player.Losses)
	return err
}

func (d *Database) GetPlayers() ([]domain.Player, error) {
	rows, err := d.db.Query("SELECT id, name, level, wins, losses FROM players WHERE deleted = FALSE")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []domain.Player
	for rows.Next() {
		var p domain.Player
		if err := rows.Scan(&p.ID, &p.Name, &p.Level, &p.Wins, &p.Losses); err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, nil
}

func (d *Database) SaveTournament(t domain.Tournament) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt := `INSERT INTO tournaments (id, name, winner_id) VALUES (?, ?, ?)`
	var winnerID *int
	if t.Winner.ID != 0 {
		winnerID = &t.Winner.ID
	}

	_, err = tx.Exec(stmt, t.ID, t.Name, winnerID)
	if err != nil {
		return err
	}

	for _, player := range t.Players {
		_, err = tx.Exec(`
            INSERT INTO tournament_players (tournament_id, player_id) 
            VALUES (?, ?)`, t.ID, player.ID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (d *Database) AddPlayerToTournament(tournamentID int, playerID int) error {
	stmt := `INSERT INTO tournament_players (tournament_id, player_id) 
             VALUES (?, ?)`

	_, err := d.db.Exec(stmt, tournamentID, playerID)
	return err
}

func (d *Database) GetTournamentPlayers(tournamentID int) ([]domain.Player, error) {
	query := `
        SELECT p.id, p.name, p.level, p.wins, p.losses 
        FROM players p
        JOIN tournament_players tp ON p.id = tp.player_id
        WHERE tp.tournament_id = ? 
        AND p.deleted = FALSE 
        AND tp.deleted = FALSE`

	rows, err := d.db.Query(query, tournamentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []domain.Player
	for rows.Next() {
		var p domain.Player
		if err := rows.Scan(&p.ID, &p.Name, &p.Level, &p.Wins, &p.Losses); err != nil {
			return nil, err
		}
		players = append(players, p)
	}
	return players, nil
}

func (d *Database) GetPlayerTournaments(playerID int) ([]domain.Tournament, error) {
	query := `
        SELECT t.id, t.name, t.winner_id 
        FROM tournaments t
        JOIN tournament_players tp ON t.id = tp.tournament_id
        WHERE tp.player_id = ? 
        AND t.deleted = FALSE 
        AND tp.deleted = FALSE`

	rows, err := d.db.Query(query, playerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tournaments []domain.Tournament
	for rows.Next() {
		var t domain.Tournament
		var winnerID sql.NullInt64
		if err := rows.Scan(&t.ID, &t.Name, &winnerID); err != nil {
			return nil, err
		}
		tournaments = append(tournaments, t)
	}
	return tournaments, nil
}

func (d *Database) DeletePlayer(playerID int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE tournament_players SET deleted = TRUE WHERE player_id = ?`, playerID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE players SET deleted = TRUE WHERE id = ?`, playerID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Database) DeleteTournament(tournamentID int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE tournament_players SET deleted = TRUE WHERE tournament_id = ?`, tournamentID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE tournaments SET deleted = TRUE WHERE id = ?`, tournamentID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Database) RemovePlayerFromTournament(tournamentID int, playerID int) error {
	stmt := `UPDATE tournament_players SET deleted = TRUE 
             WHERE tournament_id = ? AND player_id = ?`

	_, err := d.db.Exec(stmt, tournamentID, playerID)
	return err
}

func (d *Database) RestorePlayer(playerID int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE players SET deleted = FALSE WHERE id = ?`, playerID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE tournament_players SET deleted = FALSE WHERE player_id = ?`, playerID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (d *Database) RestoreTournament(tournamentID int) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`UPDATE tournaments SET deleted = FALSE WHERE id = ?`, tournamentID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE tournament_players SET deleted = FALSE WHERE tournament_id = ?`, tournamentID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
