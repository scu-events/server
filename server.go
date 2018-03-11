package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = <ADD HOSTNAME> //probably localhost
	port     = <ADD PORNUMBER> //probably 5432
	user     = <ADD USERNAME>
	password = <ADD PASSWORD>
	dbname   = <ADD DB NAME>
)

var (
	summary  string
	location string
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ //Access Server
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo) //Validate credentials
	if err != nil {
		panic(err)
	}

	err = db.Ping() //Open Connection
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT summary, location FROM data WHERE location = $1", "Benson") //Add different parameters here based on user input
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //Iterate over rows, scanning data we want
	for rows.Next() {
		err := rows.Scan(&summary, &location)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(summary, location) //print out the results <-- we can change to put into array
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
