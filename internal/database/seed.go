package database

import (
	"fmt"
)

func (d *Database) SeedDatabase() error {
	var count int
	err := d.db.QueryRow("SELECT COUNT(*) FROM players").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		fmt.Println("Database already seeded, skipping...")
		return nil
	}

	players := []struct {
		name   string
		level  int
		wins   int
		losses int
	}{
		{"João \"JohnPro\" Silva", 8, 25, 12},
		{"Maria \"M4st3r\" Santos", 7, 20, 15},
		{"Pedro \"PkMaster\" Oliveira", 5, 15, 18},
		{"Ana \"Destroyer\" Costa", 9, 28, 8},
		{"Carlos \"Rush\" Ferreira", 3, 10, 22},
		{"Julia \"Queen\" Lima", 6, 18, 16},
		{"Lucas \"LuckyStar\" Martins", 4, 12, 20},
		{"Isabella \"IceQueen\" Souza", 10, 30, 5},
		{"Miguel \"TheBoss\" Pereira", 7, 22, 14},
		{"Sofia \"SwiftBlade\" Alves", 5, 16, 19},
		{"Gabriel \"Phoenix\" Rodrigues", 8, 24, 11},
		{"Laura \"Lightning\" Carvalho", 6, 19, 17},
		{"Rafael \"RedBaron\" Ribeiro", 4, 13, 21},
		{"Beatriz \"BattleQueen\" Gomes", 9, 27, 9},
		{"Guilherme \"Ghost\" Fernandes", 2, 8, 25},
		{"Mariana \"Mage\" Barbosa", 7, 21, 13},
		{"Felipe \"Flame\" Castro", 5, 17, 18},
		{"Carolina \"Chaos\" Cardoso", 8, 23, 10},
		{"Daniel \"Dragon\" Rocha", 6, 19, 16},
		{"Amanda \"Angel\" Santos", 3, 11, 23},
	}

	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`
        INSERT INTO players (name, level, wins, losses, deleted) 
        VALUES (?, ?, ?, ?, FALSE)
    `)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, p := range players {
		_, err = stmt.Exec(p.name, p.level, p.wins, p.losses)
		if err != nil {
			return fmt.Errorf("error inserting player %s: %v", p.name, err)
		}
		fmt.Printf("Inserted player: %s\n", p.name)
	}

	return tx.Commit()
}
