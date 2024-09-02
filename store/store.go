package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kylemartinelli/surflineData/types"
)


func PrepareDataDb (r types.Main, url string, db *sql.DB) {
	insertIntoSpots(r, url, db)


	if len(r.Spot.TravelDetails.BreakType) == 0 && len(r.Spot.TravelDetails.BoardTypes) == 0 {
		return
	}


}

func insertIntoSpots(r types.Main, url string, db *sql.DB)  {

	surflineId := r.Spot.ID
	surflineUrl := url
	spotName := r.Spot.Name
	lat := r.Spot.Lat
	lon := r.Spot.Lon

	insrtStmnt := `INSERT INTO "spots" ("surflineid", "surflineurl", "spotname", "lat", "lon") VALUES ($1, $2, $3, $4, $5)`

	res, err := db.Exec(insrtStmnt, surflineId, surflineUrl, spotName, lat, lon)
	if err != nil {
		log.Fatal("failed to insert into spots", err)
	} else {
		fmt.Println(res)
	}

}