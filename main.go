package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	fmt.Println("connecting")
	// these details match the docker-compose.yml file.
	postgresInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		"localhost", 5432, "postgres", "Namle311", "user")
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v", err))
	}
	defer db.Close()
	log.Println("Connected to database")

}
