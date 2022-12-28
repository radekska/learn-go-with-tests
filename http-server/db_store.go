package poker

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewPostgresPlayerStore(db *sql.DB) *PostgresPlayerStore {
	return &PostgresPlayerStore{db: db}
}

type PostgresPlayerStore struct {
	mu sync.Mutex
	db *sql.DB
}

func getPlayerScore(db *sql.DB, name string) (int, error) {
	row := db.QueryRow("SELECT score FROM players WHERE name = $1", name)
	if err := row.Err(); err != nil {
		return 0, err
	}
	var score int
	if err := row.Scan(&score); err != nil {
		return 0, err
	}
	return score, nil
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	score, err := getPlayerScore(p.db, name)
	if err != nil {
		return 0
	}
	return score
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	p.mu.Lock()

	score, err := getPlayerScore(p.db, name)
	if err != nil {
		_, err = p.db.Exec("INSERT INTO players VALUES($1, $2)", name, 1)
		if err != nil {
			return
		}
	} else {
		_, err = p.db.Exec("UPDATE players SET score=$1+1 WHERE name=$2", score, name)
		if err != nil {
			return
		}
	}
	p.mu.Unlock()
}

func (p *PostgresPlayerStore) GetLeague() League {
	row, err := p.db.Query("SELECT name, score FROM players")
	if err != nil {
		panic(err)
	}
	var players League
	for row.Next() {
		var name string
		var score int

		if err := row.Scan(&name, &score); err != nil {
			panic(err)
		}
		players = append(players, Player{Name: name, Wins: score})
	}

	players.Sort()
	return players
}
