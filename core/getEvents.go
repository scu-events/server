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

	var events Events

	for _, item := range from_calendar.Items {
		fmt.Printf("%+v\n", item)
		events = append(events, Event{
			StartDateTime: string(item.Start.DateTime),
			EndDateTime:   string(item.End.DateTime),
			Summary:       string(item.Summary),
			Title:         string(item.Summary),
			HTMLLink:      string(item.HtmlLink),
			Location:      string(item.Location),
			Tags:          []string{},
		})
	}

	data, err := json.Marshal(ReturningData{Data: events})
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Write(data)
}
