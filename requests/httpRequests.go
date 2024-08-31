package requests

import (
	"fmt"
	"net/url"
)

func PingSurflineServices(id string) {
	baseUrl, err := url.Parse("https://services.surfline.com/kbyg/spots/reports")
	if err != nil {
		panic(err)
	}

	queryParams := url.Values{}
	queryParams.Add("spotId", id)

	baseUrl.RawQuery = queryParams.Encode()

	fmt.Println(baseUrl.String())
}