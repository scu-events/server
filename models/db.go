package models

//golang sql package and psql driver
import (
	"database/sql"

	_ "github.com/lib/pq"
)

//NewDB : Function to communicate with db
func NewDB(dataSourceName string) (*sql.DB, error) {

	//Validates arguments but doesn't create connection
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	//Creates/Checks connection
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
