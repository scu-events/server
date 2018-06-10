package main

//import packages, globals,
import (
	"net/http"
	"server/calendar"
)

func main() {

	http.HandleFunc("/events", gethandler) //may need to replace /events with whatever url we are passing
	http.ListenAndServe(":3000", nil)      //i think google calendar api might like it a certain way to parse
}

//HTTP GET Handler
func gethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	calendar.GetData(w, r)

}
