package core

import (
	"fmt"
	"google.golang.org/api/calendar/v3"
	"log"
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
	calendarID := "primary"

	for _, event := range events {
		// need to deal with ExtendedProperties
		// item := &calendar.Event{
		// Summary:     event.Title,
		// Location:    event.Location,
		// Description: event.Summary,
		// Start: &calendar.EventDateTime{
		// DateTime: sT,
		// TimeZone: "America/Los Angeles",
		// },
		// End: &calendar.EventDateTime{
		// DateTime: eT,
		// TimeZone: "America/Los Angeles",
		// },
		// // ExtendedProperties: &calendar.EventExtendedProperties{
		// // Private: map[string]string{},
		// // },
		// }
		_, err := calendarService.Events.Insert(calendarID, &event).Do()
		if err != nil {
			log.Fatalf("Unable to create event. %v\n", err)
		}
	}
}
