package requests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/kylemartinelli/surflineData/store"
	"github.com/kylemartinelli/surflineData/types"
)



func PingSurflineServices(ids [][]string, db *sql.DB)  {

	baseUrl, err := url.Parse("https://services.surfline.com/kbyg/spots/reports")
	if err != nil {
		panic(err)
	}

	idLen := len(ids)
	for i := 0; i < idLen; i++ {
		//currId := "5842041f4e65fad6a7708813"
		fmt.Println(ids[i], ids[i][1])
		currId := ids[i][1]
		queryParams := url.Values{}
		queryParams.Add("spotId", currId)
		baseUrl.RawQuery = queryParams.Encode()

	 response := makeRequest(baseUrl.String())


	  store.PrepareDataDb(response, baseUrl.String(), db, currId)


	}




}

func makeRequest (url string) (types.Main) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("failed to get from surfline services: ", err)
	}
	defer response.Body.Close()

	var responseData types.Main

	body , err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal("Failed to read response body: ", err)
	}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Fatal("Failed to unmarshal json: ", err)
	}



	return responseData


}