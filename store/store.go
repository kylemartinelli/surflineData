package store

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kylemartinelli/surflineData/types"
	"github.com/lib/pq"
)


func PrepareDataDb (r types.Main, url string, db *sql.DB, surflineId string) {
	spotId := surflineId
	  //insertIntoSpots(r, url, db)
	  insertIntoBreadcrumbs(r.Spot.Breadcrumb, db, spotId)


	if len(r.Spot.TravelDetails.BreakType) == 0 && len(r.Spot.TravelDetails.BoardTypes) == 0 {
		return
	}
	  insertSpotChar(r.Spot.TravelDetails, db, spotId)
	 insertIdealCon(r.Spot.TravelDetails.Best, db, spotId)


}

func insertIntoSpots(r types.Main, url string, db *sql.DB)  {

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

}

func insertIntoBreadcrumbs(r []types.Breadcrumb , db *sql.DB, spotID string) {



	crumbMap := map[string]string{"countryName": "", "countryUrl": "", "regionName": "", "regionUrl": "", "subRegionName": "", "subRegionUrl": "", "areaName": "", "areaUrl": ""}

	for i := 0; i < len(r); i++ {
		switch i {
			case 0 :
				crumbMap["countryName"] = r[0].Name
				crumbMap["countryUrl"] = r[0].Href
			 case 1 :
				crumbMap["regionName"] = r[1].Name
				crumbMap["regionUrl"] = r[1].Href
				 case 2:
				crumbMap["subRegionName"] = r[2].Name
				crumbMap["subRegionUrl"] = r[2].Href
				 case  3:
				crumbMap["areaName"] = r[3].Name
				crumbMap["areaUrl"] = r[3].Href
			}


	}


	insrtStmt := `INSERT INTO "spotbreadcrumbs" ("spotid", "country", "countryurl", "region", "regionurl", "subregion", "subregionurl", "area", "areaurl") VALUES ((SELECT "id" FROM spots WHERE "surflineid" = $1 ), $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := db.Exec(insrtStmt, spotID, crumbMap["countryName"], crumbMap["countryUrl"], crumbMap["regionName"], crumbMap["regionUrl"], crumbMap["subRegionName"], crumbMap["subRegionUrl"], crumbMap["areaName"], crumbMap["areaUrl"])
	if err != nil {
		log.Fatal("failed to insert into breadcrumbs: ", err, "spotId: ", spotID)
	}


}

func insertSpotChar(r types.TravelDetails, db *sql.DB, spotId string) {


	insrtStmnt := `INSERT INTO "spotcharacteristics" ("spotid", "crowdfactordescription", "crowdfactorrating", "crowdfactorsummary", "localvibedescription", "localviberating", "localvibesummary", "shoulderburndescription", "shoulderburnrating", "shoulderburnsummary", "spotratingdescription", "spotratingrating","spotratingsummary", "waterqualitydescription", "waterqualityrating", "waterqualitysummary","abilitydescription", "bottomtypedescription", "spotdescription", "abilitylevels", "bottomtype", "boardtypes", "breaktype", "hazards" ) VALUES((SELECT "id" FROM spots WHERE "surflineid" = $1 ), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24)`
	_, err := db.Exec(insrtStmnt, spotId, r.CrowdFactor.Description, r.CrowdFactor.Rating, r.CrowdFactor.Summary, r.LocalVibe.Description, r.LocalVibe.Rating, r.LocalVibe.Summary, r.ShoulderBurn.Description, r.ShoulderBurn.Rating, r.ShoulderBurn.Summary, r.SpotRating.Description, r.SpotRating.Rating, r.SpotRating.Summary, r.WaterQuality.Description, r.WaterQuality.Rating, r.WaterQuality.Summary, r.AbilityLevels.Description, r.Bottom.Description, r.Description, pq.Array(r.AbilityLevels.Levels), pq.Array(r.Bottom.Value), pq.Array(r.BoardTypes), pq.Array(r.BreakType), r.Hazards)
	if err != nil {
		log.Fatal("Failed to insert into spot char: ",err, "spotId: " ,spotId)
	}

}

func insertIdealCon (r types.Best, db *sql.DB, spotId string) {
	insrtStmnt := `INSERT INTO "idealconditions" ("spotid", "idealseason", "idealseasondescription", "idealtide", "idealtidedescription", "idealsize", "idealsizedescription", "idealwinddirections", "idealwinddirectionsdescription", "idealswelldirections", "idealswelldirectionsdescription") VALUES((SELECT "id" FROM spots WHERE "surflineid" = $1 ), $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`
	_, err := db.Exec(insrtStmnt, spotId, pq.Array(r.Season.Value), r.Season.Description, pq.Array(r.Tide.Value),r.Tide.Description, pq.Array(r.Size.Value), r.Size.Description, pq.Array(r.WindDirection.Value), r.WindDirection.Description, pq.Array(r.SwellDirection.Value), r.SwellDirection.Description)
	if err != nil {
		log.Fatal("failed to insert into ideal con: ", err, "sportId: ", spotId)
	}

}
