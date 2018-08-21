package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed"

	"google.golang.org/api/calendar/v3"
)

func CrawlGoogleCal(url string, calendarId string) ([]calendar.Event, error) {
	var res []calendar.Event

	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	var data map[string][]calendar.Event
	json.Unmarshal(body, &data)

	for _, item := range data["items"] {
		// maybe define the exact structure we are interested in
		res = append(res, calendar.Event{
			Description: item.Description,
			Start:       item.Start,
			End:         item.End,
			Location:    item.Location,
			Summary:     item.Summary,
		})
	}
	return res, nil
}

func CrawlOfficialCal(url string) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	fmt.Println(feed.Title)
	fmt.Println()

	for _, element := range feed.Items {
		fmt.Println(element.Title)

		fmt.Printf("	")
		fmt.Printf(element.Description)
	}

	return
}

//Using natural language processing to determine if an event has free food
func freeFood(description string) bool {
	foodWords := [7]string{"food", "refreshments", "dinner", "lunch", "breakfast", "snacks", "drinks"}

	for _, element := range foodWords {
		if strings.Contains(description, element) {
			return true
		}
	}

	return false
}
