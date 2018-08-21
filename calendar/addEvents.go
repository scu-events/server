package calendar

import (
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

func insertEvent(title string, location string, description string, sT string, eT string, food string, dpt string) {
	b, err := ioutil.ReadFile("client_secret.json")
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
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
			Private: map[string]string{},
		},
	}
	event.ExtendedProperties.Private["freeFood"] = food
	event.ExtendedProperties.Private["department"] = dpt
	calendarID := "primary" //this adds event to the calendar of the logged in user, we can change to url
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
}
