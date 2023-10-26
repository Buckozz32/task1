package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
	goose "github.com/pressly/goose"
)

type EnrichedMessage struct {
	ID       int    `json:"id"`
	Message  string `json:"message"`
	Enriched bool   `json:"enriched"`
}

func main() {
	dbName := "postgres"
	dbUser := "postgres"
	dbPassword := "12345"
	dbHost := "localhost"
	dbPort := 5432

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}
	if err := goose.Run("up", db, "./migrations"); err != nil {
		panic(err)
	}

	message := EnrichedMessage{
		ID:       1,
		Message:  "Hello, world!",
		Enriched: true,
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("", string(jsonData))
	if err != nil {
		panic(err)
	}

	db.Close()
}
