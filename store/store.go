package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kylemartinelli/surflineData/types"
)


func PrepareDataDb (r types.Main, url string, db *sql.DB) {

	spotId := insertIntoSpots(r, url, db)
	insertIntoBreadcrumbs(r.Spot.Breadcrumb, db, spotId)


	if len(r.Spot.TravelDetails.BreakType) == 0 && len(r.Spot.TravelDetails.BoardTypes) == 0 {
		return
	}


}

func insertIntoSpots(r types.Main, url string, db *sql.DB) string {

	surflineId := r.Spot.ID
	surflineUrl := url
	spotName := r.Spot.Name
	lat := r.Spot.Lat
	lon := r.Spot.Lon

	insrtStmnt := `INSERT INTO "spots" ("surflineid", "surflineurl", "spotname", "lat", "lon") VALUES ($1, $2, $3, $4, $5)`

	res, err := db.Exec(insrtStmnt, surflineId, surflineUrl, spotName, lat, lon)
	fmt.Println(res)
	if err != nil {
		log.Fatal("failed to insert into spots", err)
	}
	return surflineId
}

func insertIntoBreadcrumbs(r []types.Breadcrumb , db *sql.DB, spotID string) {

	countryName := r[0].Name
	countryUrl := r[0].Href
	regionName := r[1].Name
	regionUrl := r[1].Href
	subregionName := r[2].Name
	subregionUrl := r[2].Href

	if len(r) == 4 {
		areaName := r[3].Name
		areaUrl := r[3].Href

		insrtStmt := `INSERT INTO "spotbreadcrumbs" ("spotid", "country", "countryurl", "region", "regionurl", "subregion", "subregionurl", "area", "areaurl") VALUES ((SELECT "id" FROM spots WHERE "surflineid" = $1 ), $2, $3, $4, $5, $6, $7, $8, $9)`
		_, err := db.Exec(insrtStmt, spotID, countryName, countryUrl, regionName, regionUrl, subregionName, subregionUrl, areaName, areaUrl)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	insrtStmt := `INSERT INTO "spotbreadcrumbs" ("spotid", "country", "countryurl", "region", "regionurl", "subregion", "subregionurl") VALUES ((SELECT "id" FROM spots WHERE "surflineid" = $1 ), $2, $3, $4, $5, $6, $7)`
	_, err := db.Exec(insrtStmt, spotID, countryName, countryUrl, regionName, regionUrl, subregionName, subregionUrl)
	if err != nil {
		log.Fatal(err)
	}


}

func insertChar() {

}
