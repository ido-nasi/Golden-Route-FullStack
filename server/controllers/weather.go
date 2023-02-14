package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Struct to parse the weather api response into
type WeatherApiResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hourly    Hourly  `json:"hourly"`
	Error     bool    `json:"error"`
	Reason    string  `json:"reason"`
}

type Hourly struct {
	Time        []string  `json:"time"`
	Temperature []float64 `json:"temperature_2m"`
}

// Struct function to unmarshal the response into the WeatherApiResponse struct
func (W *WeatherApiResponse) UnmarshalResponse(response *http.Response) {

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal([]byte(responseBytes), &W)
	if err != nil {
		log.Println(err)
	}
}

// Struct to parse the frontend request parameters into
type RequestParams struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Date      string `json:"date"`
}
