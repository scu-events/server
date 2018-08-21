package main

import (
	"./core"
	"net/http"
)

func main() {
	/*
		calendarId := "info.scuevents@gmail.com"
		//GOOGLE CAL URLS
		//acm club: http://acm.engr.scu.edu/events
		var url_google_acm = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"

		//offical calendar of CSO's and RSO's: https://www.scu.edu/csi/calendar/
		//	var url_google_clubs = "https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs"

		//	var officialUrl = "https://lwcal.scu.edu/live/rss/events/exclude_group/Math%20Tutors/exclude_tag/private/header/All%20Events"

		//CrawlGoogleCal(url_google_clubs)
		events, _ := core.CrawlGoogleCal(url_google_acm, calendarId)
		core.AddEvents(events)
	*/

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
