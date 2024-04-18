package api_test

import (
	"getWeather/api"
	"testing"
)

func TestGetWeather(t *testing.T) {
	_, err := api.GetWeather(api.Coord{0, 0})

	if err == nil {
		t.Error("Expected error when calling GetWeather() with 0 latitude and 0 longitude")
	}
}
