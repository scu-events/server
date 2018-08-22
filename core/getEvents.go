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

	m, ok := r.URL.Query()["month"]
	var month time.Month
	if !(ok && len(m) >= 0) {
		month = time.Now().Month()
	} else {
		month = months[m[0]]
	}

	y, ok := r.URL.Query()["year"]
	var year int
	if !(ok && len(y) >= 0) {
		year = time.Now().Year()
	} else {
		ye, err := strconv.Atoi(y[0])
		year = ye
		if err != nil {
			year = time.Now().Year()
		}
	}

	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		loc = time.UTC
	}

	calendar_service := GetCalendarService()
	t := time.Date(year, month, 0, 0, 0, 0, 0, loc)
	from_calendar, err := calendar_service.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t.Format(time.RFC3339)).TimeMax(t.AddDate(0, 1, 0).Format(time.RFC3339)).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next five of the user's events: %v", err)
	}

	data, err := json.Marshal(ReturningData{Data: from_calendar.Items})
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Write(data)
}
