package main

import (
	"./core"
	"google.golang.org/api/calendar/v3"
	"os"
	"fmt"
	"sync"
)

func main() {
	args := os.Args[1:]

	// TODO: In the future, we should load all the calendars from a text/JSON file online --Matthew
	//GOOGLE CAL URLS
	google_cals := make(map[string]string)
	//acm club: http://acm.engr.scu.edu/events
	google_cals["acm"] = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	//Official CSO/RSO Calendar
	google_cals["cso"] = "https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs"

	var wg sync.WaitGroup
	wg.Add(len(args))
	events_ch := make(chan []calendar.Event)

	for _, arg := range args {
		go func(url string) {
			defer wg.Done()

			events, err := core.CrawlGoogleCal(url)
			if err == nil {
				wg.Add(len(args))
				events_ch <- events
			} else { 
				panic(err)
			}
		}(google_cals[arg])
	}

	go func() {
		for events := range events_ch {
			defer wg.Done()
			fmt.Printf("%d", len(events))
			core.AddEvents(events)
		}
	}()

	wg.Wait()
}
