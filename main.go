package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/stdlib"
)

func main() {
	err := CreateExtensions()
	if err != nil {
		fmt.Printf("Error creating extensions: %v", err)
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9001"
	}
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":"+port, nil)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wow much server")
}

// CreateExtensions creates some Postgres extensions, and returns errors if they fail
func CreateExtensions() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP EXTENSION IF EXISTS postgis")
	_, err = db.Exec("CREATE EXTENSION postgis")
	if err != nil {
		return err
	}
	return nil
}
