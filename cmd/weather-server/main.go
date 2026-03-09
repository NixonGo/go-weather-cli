package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-weather-cli/internal/weather"
	"html/template"
	"net/http"
)

type PageData struct {
	Weather weather.WeatherResponse
	Image   string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

var httpRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "weather_http_requests_total",
		Help: "Total number of HTTP requests",
	},
	[]string{"path"},
)

func handler(w http.ResponseWriter, r *http.Request) {
	httpRequests.WithLabelValues(r.URL.Path).Inc()
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
	prometheus.MustRegister(httpRequests)

	http.HandleFunc("/", handler)
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}
