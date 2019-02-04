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
	// GOOGLE CAL URLS
	google_cals := make(map[string]string)
	// acm club: http://acm.engr.scu.edu/events
	google_cals["acm"] = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	// Official CSO/RSO Calendar
	google_cals["cso"] = "https://www.googleapis.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&&timeMin=2018-02-24T00:00:00Z&timeMax=2099-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	// Engineering Student Events Calendar - https://www.scu.edu/engineering/beyond-the-classroom/student-organizations/
	google_cals["engr"] = "https://www.googleapis.com/calendar/v3/calendars/scu.edu_5m30rgamr2hc8ch6e54at49suo@group.calendar.google.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"

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
