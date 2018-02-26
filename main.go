package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"

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
	_, err = db.Exec("DROP EXTENSION IF EXISTS pg_stat_statements")
	_, err = db.Exec("CREATE EXTENSION pg_stat_statements")
	if err != nil {
		return err
	}
	return nil
}

// PokeRedis pokes redis, and returns an error if it is unsuccessful
func PokeRedis() error {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	fmt.Println(pong, err)
	return nil
}
