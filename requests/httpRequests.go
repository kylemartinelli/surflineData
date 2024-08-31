package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"

	"github.com/lylemartinelli/surflineData/types"
)



func PingSurflineServices(ids []string)  {

	baseUrl, err := url.Parse("https://services.surfline.com/kbyg/spots/reports")
	if err != nil {
		panic(err)
	}

	for range ids {
		queryParams := url.Values{}
		queryParams.Add("spotId", "5842041f4e65fad6a7708813")
		baseUrl.RawQuery = queryParams.Encode()

		makeRequest(baseUrl.String())
		//fmt.Println(data)
		break;
	}




}

func makeRequest (url string) string {
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
fmt.Println(responseData.Spot.Breadcrumb[1])

	return string(body)


}