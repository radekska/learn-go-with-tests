package poker

import (
	"database/sql"
	"fmt"
	"net/url"
)

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
