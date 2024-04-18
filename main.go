package main

import (
	"fmt"
	"getWeather/api"
)

func main() {
	fmt.Println("Hello, World!")
	weather, err := api.GetWeather()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Temperature: %v\n", weather.Temperature)
}
