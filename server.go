package main

import (
	"./calendar"
	"net/http"
)

func main() {
	http.HandleFunc("/api/events", withCommonConfig(gethandler))
	http.ListenAndServe(":4000", nil)
}

//HTTP GET Handler
func gethandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	calendar.GetData(w, r)
}

func withCommonConfig(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	}
}
