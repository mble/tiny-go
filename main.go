package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	fmt.Println("Testing extension creation")
	err := CreateExtensions()
	if err != nil {
		log.Fatalf("Error creating extensions: %v", err)
	}
	fmt.Println("Tested extensions")
}

// CreateExtensions creates some Postgres extensions, and returns errors if they fail
func CreateExtensions() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	_, err = db.Exec("CREATE EXTENSION postgis")
	if err != nil {
		return err
	}
	return nil
}
