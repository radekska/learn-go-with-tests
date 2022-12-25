package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	"net/url"
	"sync"
)

func NewPostgresPlayerStore() *PostgresPlayerStore {
	return &PostgresPlayerStore{}
}

type PostgresPlayerStore struct {
	mu sync.Mutex
}

func getPlayerScore(db *sql.DB, name string) (int, error) {
	row := db.QueryRow("SELECT score FROM players WHERE name = $1", name)
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRow", err)
		return 0, err
	}
	var score int
	if err := row.Scan(&score); err != nil {
		fmt.Println("row.Scan", err)
		return 0, err
	}
	return score, nil
}

func GetDatabase() (*sql.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "players",
	}
	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		return nil, err
	}
	return db, nil
}

func (p *PostgresPlayerStore) GetPlayerScore(name string) int {
	db, err := GetDatabase()
	defer db.Close()

	score, err := getPlayerScore(db, name)
	if err != nil {
		fmt.Printf("failed to retrieve score for %s player\n", name)
		return 0
	}
	return score
}

func (p *PostgresPlayerStore) RecordWin(name string) {
	p.mu.Lock()

	db, err := GetDatabase()
	defer db.Close()

	score, err := getPlayerScore(db, name)
	if err != nil {
		_, err = db.Exec("INSERT INTO players VALUES($1, $2)", name, 1)
		if err != nil {
			fmt.Println("db.Exec", err)
			return
		}
	} else {
		_, err = db.Exec("UPDATE players SET score=$1+1 WHERE name=$2", score, name)
		if err != nil {
			fmt.Println("db.Exec", err)
			return
		}
	}
	p.mu.Unlock()
}
