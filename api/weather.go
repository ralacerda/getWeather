package api

import (
	"encoding/json"
	"fmt"
	"getWeather/datatype"
	"io"
	"net/http"
)

const url = "https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m"

type Coord struct {
	Latitude  float64
	Longitude float64
}

func GetWeather(coord Coord) (*datatype.Weather, error) {
	if coord.Latitude == 0 && coord.Longitude == 0 {
		return nil, fmt.Errorf("invalid coordinates")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var response OpenWeatherResponse
	json.Unmarshal(bodyBytes, &response)

	weather := &datatype.Weather{Temperature: response.Current.Temperature2M}
	return weather, nil

}
