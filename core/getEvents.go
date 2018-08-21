package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//GetData function called in main
func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	calendar_service := GetCalendarService()
	t := time.Now().Format(time.RFC3339)
	from_calendar, err := calendar_service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(5).OrderBy("startTime").Do() //change from primary
	if err != nil {
		log.Fatalf("Unable to retrieve next five of the user's events: %v", err)
	}

	data, err := json.Marshal(ReturningData{Data: from_calendar.Items})
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Write(data)
}
