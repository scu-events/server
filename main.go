package main

//import packages, globals,
import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"server/models"
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
	db, err := models.NewDB("postgres://postgres:" + *pwd + "@localhost/events?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db: db}

	http.HandleFunc("/events", env.gethandler)
	http.ListenAndServe(":3000", nil)
}

//HTTP GET Handler
func (env *Env) gethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	events, err := models.OutputJSON(env.db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//Print out results
	fmt.Println(events)
}
