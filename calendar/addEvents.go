package calendar

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func insertEvent(title string, location string, description string, sT string, eT string, food string, depart string) {
	b, err := ioutil.ReadFile("client_secret.json")
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	srv, err := calendar.New(getClient(config))
	event := &calendar.Event{
		Summary:     title,
		Location:    location,
		Description: description,
		Start: &calendar.EventDateTime{
			DateTime: sT,
			TimeZone: "America/Los Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: eT,
			TimeZone: "America/Los Angeles",
		},
		ExtendedProperties: &calendar.EventExtendedProperties{
			Private: make(map[string]string),
		},
	}
	calendarID := "primary" //this adds event to the calendar of the logged in user, we can change to url
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}

/*//Seperate function to add extended properties with patching, rather than insert (will handle Nulls better)
func ExtendedProperties(food bool, depart string) {
	b, err := ioutil.ReadFile("client_secret.json")
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	srv, err := calendar.New(getClient(config))
	event := &calendar.Event{
		"extendedProperties": {
			"private": {
				freeFood:   food,
				department: depart,
			},
		},
	}
	calendarId := "primary" //this adds to calendar of the logged in user
	event, err = srv.Events.Patch(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Error with properties. %v\n", err)
	}
	fmt.Printf("Extended Properties added to: %s\n", event.HtmlLink)
}
*/
