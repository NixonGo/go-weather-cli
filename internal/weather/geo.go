package weather

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

func GetCityData(city string) (Geo, error) {

	params := url.Values{
		"q":     {city},
		"appid": {apiKey},
		"limit": {"5"},
	}

	link := "http://api.openweathermap.org/geo/1.0/direct?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return Geo{}, errors.New("error while creating request to " + link)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return Geo{}, errors.New("error: " + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Geo{}, errors.New("error: " + resp.Status)
	}

	var geo []Geo

	if err = json.NewDecoder(resp.Body).Decode(&geo); err != nil {
		return Geo{}, errors.New("error decoding response: " + err.Error())
	}

	if len(geo) == 0 {
		return Geo{}, errors.New("no weather data found")
	}

	return geo[0], nil
}
