package main

import (
	"fmt"
	"getWeather/api"
)

func main() {
	fmt.Println("Hello, World!")
	weather, err := api.GetWeather(api.Coord{Latitude: 52.52, Longitude: 13.41})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Temperature: %v\n", weather.Temperature)
}
