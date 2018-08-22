package main

import (
	"./core"
	"net/http"
)

func main() {
	http.HandleFunc("/api/events", withCommonConfig(core.GetData))
	http.ListenAndServe(":4000", nil)
}

func withCommonConfig(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	}
}
