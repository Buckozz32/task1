package main

import (
 "database/sql"
 "fmt"
 "log"
 

 "github.com/pressly/goose/v3"
 _ "github.com/lib/pq"
)

func ConncetDataBase() {
 
 dbConfig := "user=postgres password=12345d bname=postgres sslmode=disable"

 
 migrationsDir := "path/to/migrations"

 
 db, err := sql.Open("postgres", dbConfig)
 if err != nil {
  log.Fatal(err)
 }

 
 goose.SetDialect("postgres")
 goose.SetTableName("goose_db_version")

 
 err = goose.Up(db, migrationsDir)
 if err != nil {
  log.Fatal(err)
 }

 err = db.Close()
 if err != nil {
  log.Fatal(err)
 }

 fmt.Println("Миграции успешно применены!")
}
