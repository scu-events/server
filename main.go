package main

//import packages, globals,
import (
	"fmt"
	"net/http"
	"server/calendar"
)

func main() {

	http.HandleFunc("/events", gethandler)
	http.ListenAndServe(":3000", nil)
}

//HTTP GET Handler
func gethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	events, err := calendar.GetData(r)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//Print out results
	fmt.Println(events)
}
