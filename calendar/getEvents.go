package calendar

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

type eventData struct {
	StartDateTime string `json:"start_date_time"`
	EndDateTime   string `json:"end_date_time"`
	Location      string `json:"location"`
	HTMLLink      string `json:"html_link"`
	Summary       string `json:"summary"`
	Title         string `json:"title"`
}

type eventsData []eventData

type returnedDate struct {
	Data interface{} `json:"data"`
}

//GetData function called in main
func GetData(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved client_secret.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := calendar.New(getClient(config))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(5).OrderBy("startTime").Do() //change from primary
	if err != nil {
		log.Fatalf("Unable to retrieve next five of the user's events: %v", err)
	}

	var acc eventsData

	for _, item := range events.Items {
		fmt.Printf("%+v\n", item)
		acc = append(acc, eventData{
			StartDateTime: string(item.Start.DateTime),
			EndDateTime:   string(item.End.DateTime),
			Summary:       string(item.Summary),
			Title:         string(item.Summary),
			HTMLLink:      string(item.HtmlLink),
			Location:      string(item.Location),
		})
	}

	res := returnedDate{Data: acc}

	eventsJSON, err := json.Marshal(res)
	if err != nil {
		fmt.Fprintf(w, "Error: %s", err)
	}

	w.Write(eventsJSON)
}
