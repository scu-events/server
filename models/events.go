package models

import "database/sql"

//Global struct to hold variables/datatype of information we will grab
type Event struct {
	Summary     string
	Description string
	Location    string
	Starttime   string
	Endtime     string
}

//Function to query data
func AllEvents(db *sql.DB) ([]*Event, error) {

	/*I assume we are going to pass slices here
	based off of the input from the user on the front end
	to retreive specific data to display, for now just going
	to grab everything*/
	rows, err := db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	//Need to close because reasons
	defer rows.Close()

	events := make([]*Event, 0)

	//Search through rows for data
	for rows.Next() {
		event := new(Event)

		//Scan rows to grab information from columns
		err := rows.Scan(&event.Summary, &event.Description, &event.Location, &event.Starttime, &event.Endtime)
		if err != nil {
			return nil, err
		}

		//append information we just grabbed to a slice, then go to the next row
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
