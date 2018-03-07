package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/garyburd/redigo/redis"

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

// rootHandler handles the root
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wow much server")
}

// CreateExtensions creates some Postgres extensions, and returns errors if they fail
func CreateExtensions() error {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	_, err = db.Exec("DROP EXTENSION IF EXISTS pg_stat_statements")
	_, err = db.Exec("CREATE EXTENSION pg_stat_statements")
	if err != nil {
		return err
	}
	return nil
}

// PokeRedis pokes redis, and returns an error if it is unsuccessful
func PokeRedis() error {
	client, err := redis.DialURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}
	_, err = client.Do("PING")
	if err != nil {
		return err
	}
	defer client.Close()
	return nil
}
