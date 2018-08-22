package main

import (
	"./core"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	//GOOGLE CAL URLS
	google_cals := make(map[string]string)
	//acm club: http://acm.engr.scu.edu/events
	google_cals["acm"] = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	//Official CSO/RSO Calendar
	google_cals["cso"] = "https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs"

	for _, arg := range args {
		events, _ := core.CrawlGoogleCal(google_cals[arg])
		core.AddEvents(events)
	}
}
