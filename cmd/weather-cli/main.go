package main

import (
	"fmt"
	"go-weather-cli/internal/weather"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: weather <city>")
		os.Exit(1)
	}

	city := os.Args[1]

	getCityGeo, err := weather.GetCityData(city)
	if err != nil {
		fmt.Println("Ошибка получения данных о городе", err)
		os.Exit(1)
	}
	weatherNow, err := weather.СurrentWeather(getCityGeo)
	if err != nil {
		fmt.Println("Ошибка получения погоды: " + err.Error())
		os.Exit(1)
	}
	weather.PrintWeather(getCityGeo, weatherNow)
}
