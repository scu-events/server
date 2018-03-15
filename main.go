package main

//import packages, globals,
import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"scu_events/models"
)

//Global pointer to open a connection pool to db
type Env struct {
	db *sql.DB
}

func main() {

	//Take password as argument -- will change later
	var pwd = flag.String("pwd", "Incorrect Password", "server password")
	flag.Parse()

	//Function created in db.go to handle database logic
	db, err := models.NewDB("postgres://postgres:" + *pwd + "@localhost/scu_events?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db: db}

	http.HandleFunc("/events", env.eventsIndex)
	http.ListenAndServe(":3000", nil)
}

//Http handler
func (env *Env) eventsIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	events, err := models.AllEvents(env.db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//Print out results
	for _, event := range events {
		fmt.Fprintf(w, "%s, %s, %s, %s, %s\n", event.Summary, event.Description, event.Location, event.Starttime, event.Endtime)
	}
}
