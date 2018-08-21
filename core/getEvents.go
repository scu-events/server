package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var months = map[string]time.Month{
	"jan": time.January,
	"feb": time.February,
	"mar": time.March,
	"apr": time.April,
	"may": time.May,
	"jun": time.June,
	"jul": time.July,
	"aug": time.August,
	"sep": time.September,
	"oct": time.October,
	"nov": time.November,
	"dec": time.December,
}

//GetData function called in main
func GetData(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	month := r.URL.Query()["month"][0]
	y := r.URL.Query()["year"][0]
	year, err := strconv.Atoi(y)
	loc, err := time.LoadLocation("America/Los_Angeles")

	calendar_service := GetCalendarService()
	t := time.Date(year, months[month], 0, 0, 0, 0, 0, loc).Format(time.RFC3339)
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
