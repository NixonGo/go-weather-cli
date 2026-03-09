package main

import (
	"go-weather-cli/internal/weather"
	"html/template"
	"net/http"
)

type PageData struct {
	Weather weather.WeatherResponse
	Image   string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func handler(w http.ResponseWriter, r *http.Request) {

	city := r.FormValue("city")

	var page PageData

	if city != "" {

		geo, err := weather.GetCityData(city)
		if err == nil {
			wdata, err := weather.СurrentWeather(geo)
			if err == nil {

				page = PageData{
					Weather: wdata,
				}
			}
		}
	}

	tmpl.Execute(w, page)
}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
