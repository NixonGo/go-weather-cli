package weather

type Geo struct {
	Name string  `json:"name"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}
type WeatherResponse struct {
	Name       string        `json:"name"`
	Main       WeatherMain   `json:"main"`
	WeatherSky []WeatherData `json:"weather"`
}
type WeatherData struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}
type WeatherMain struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}
