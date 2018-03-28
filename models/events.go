package models

import (
	"database/sql"
	"encoding/json"
)

//Global struct to hold variables/datatype of information we will grab
type Event struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Starttime   string `json:"starttime"`
	Endtime     string `json:"endtime"`
}

func OutputJSON(db *sql.DB) (string, error) {
	rows, err := db.Query("SELECT * FROM events;")
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
