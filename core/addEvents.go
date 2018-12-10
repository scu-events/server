package core

import (
	"fmt"
	"crypto/md5"
	"google.golang.org/api/calendar/v3"
	"log"
	"encoding/hex"
	"time"
)

/*
Following works
	events := []byte(`[{"creator":{"email":"scuhackers@gmail.com","self":true},"end":{"dateTime":"2018-08-24T12:00:00-07:00"},"extendedProperties":{"private":{"freefood":"1"}},"kind":"calendar#event","organizer":{"email":"scuhackers@gmail.com","self":true},"reminders":{"useDefault":true},"start":{"dateTime":"2018-08-24T11:00:00-07:00"},"status":"confirmed","summary":"From Go Test"}]`)
	var to_insert []calendar.Event
	json.Unmarshal(events, &to_insert)
	core.AddEvents(to_insert)
*/
func AddEvents(events []calendar.Event) {
	fmt.Printf("%d events\n", len(events))

	calendarService := GetCalendarService()
	calendarID := "bo5c3s40s9cnjt2keo490744cg@group.calendar.google.com"

	for _, event := range events {		
		if(event.Start != nil && event.End != nil) {
			fmt.Printf(".")

			// Hash the title and start time to generate the event id
			// This allows duplicate events to be rejected by the google calender API
			hashInput := []byte(event.Summary + event.Start.DateTime)
			var hash = md5.Sum(hashInput)
			var hashString = hex.EncodeToString(hash[:])
			fmt.Printf(hashString)
			event.Id = hashString

			_, err := calendarService.Events.Insert(calendarID, &event).Do()
			if err != nil {
				log.Printf("Unable to create event. %v\n", err)
			}

			// Wait for 1 second after creating the event to not hit the API's limit
			// of 2,000 queries / 100 seconds 
			time.Sleep(time.Second)
		}
	}
}
