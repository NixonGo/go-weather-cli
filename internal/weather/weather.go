package weather

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func СurrentWeather(geo Geo) (WeatherResponse, error) {
	lat := strconv.FormatFloat(geo.Lat, 'f', -1, 64)
	lon := strconv.FormatFloat(geo.Lon, 'f', -1, 64)

	params := url.Values{
		"lat":   {lat},
		"lon":   {lon},
		"appid": {apiKey},
		"lang":  {"ru"},
		"units": {"metric"},
	}
	link := "https://api.openweathermap.org/data/2.5/weather?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return WeatherResponse{}, errors.New("Error while creating request to " + link)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return WeatherResponse{}, errors.New("Error: " + err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		return WeatherResponse{}, errors.New("Error: " + resp.Status)
	}
	defer resp.Body.Close()

	var weather WeatherResponse
	if err = json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return WeatherResponse{}, errors.New("Error decoding response: " + err.Error())
	}
	return weather, nil
}

func PrintWeather(city Geo, weather WeatherResponse) {
	fmt.Println("------ Погода сейчас ------\n")
	fmt.Printf("Город: %s\n", city.Name)
	fmt.Printf("Погода: %s (%s)\n", weather.WeatherSky[0].Description, weather.WeatherSky[0].Main)
	fmt.Printf("Температура: %.1f°C\n", weather.Main.Temp)
	fmt.Printf("Ощущается как: %.1f°C\n", weather.Main.FeelsLike)
	fmt.Printf("Влажность: %d%%\n", weather.Main.Humidity)
	fmt.Printf("Давление: %d hPa\n", weather.Main.Pressure)

	fmt.Println("\n---------------------")

}
