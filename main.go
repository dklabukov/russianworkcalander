package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dklabukov/russianworkcalander/types"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get(fmt.Sprintf("http://data.gov.ru/api/json/dataset/7708660670-proizvcalendar/version/20151123T183036/content?access_token=%s", os.Args[1]))

	if err != nil {
		panic(err)
	}

	dat := make([]types.Calendar, 0)

	if err = json.NewDecoder(resp.Body).Decode(&dat); err != nil {
		panic(err)
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", os.Args[2], os.Args[3], os.Args[4], os.Args[5])
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO calendar VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
ON CONFLICT ON CONSTRAINT pk_calendar 
DO UPDATE SET jan = $2, feb = $3, mar = $4, apr = $5, may = $6, jun = $7, jul = $8, aug = $9, sep = $10, oct= $11, nov = $12, dec = $13
`

	for _, c := range dat {
		_, err = db.Exec(sqlStatement, c.Year, c.Jan, c.Feb, c.Mar, c.Apr, c.May, c.Jun, c.Jul, c.Aug, c.Sep, c.Oct, c.Nov, c.Dec)
		if err != nil {
			panic(err)
		}
	}
}
