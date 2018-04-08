package models

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"
)

//Event : Global struct to hold variables/datatype of information we will grab
type Event struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Starttime   string `json:"starttime"`
	Endtime     string `json:"endtime"`
}

//OutputJSON : for HTTP Get request
func OutputJSON(db *sql.DB, r *http.Request) (string, error) {
	params, err := URLParse(r)
	rows, err := db.Query("SELECT * FROM events WHERE month = '$1', month = '$2', month = '$3, year = '$4';", params[0],  params[1],  params[2],  params[3])

	if err != nil {
		return string("Error 1"), nil
	}

	//Need to close because reasons
	defer rows.Close()

	//grabs list of column names
	columns, err := rows.Columns()
	if err = rows.Err(); err != nil {
		return string("Error 2"), nil
	}

	//count amount of columns
	count := len(columns)

	//This should preserve data types
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	return string(jsonData), nil
}

func URLParse(r *http.Request) ([4]string, error){
	var i int
	array := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	url := r.URL.String()
	parameters := strings.Split(url, "&")
	currentyear := strings.Split(parameters[1], "=")
	currentmonth := strings.Split(parameters[0], "=")
	for i=0; i<12; i++ {
		if array[i] == currentmonth[1] {
			break
		}
	}
	array2 := [4]string{array[i-1], currentmonth[1], array[i+1], currentyear[1]}
	return array2, nil
}