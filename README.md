# Go Weather App

![Go](https://img.shields.io/badge/Go-1.22-blue)
![License](https://img.shields.io/badge/license-MIT-green)
![Platform](https://img.shields.io/badge/platform-linux-lightgrey)

Simple weather application written in **Go**.

<img width="1748" height="912" alt="image" src="https://github.com/user-attachments/assets/e35bdfb0-398f-49ac-b9d0-d3d8b9471a27" />

The project provides two ways to use the weather service:

* CLI tool for terminal usage
* Web server with a simple HTML interface

Weather data is fetched from the **OpenWeather API**.

---
## Build

go build -o weather-server ./cmd/weather-server

---

# Features

* Get current weather by city
* CLI interface
* Web interface
* HTML templates
* Deployable on a VPS
* Reverse proxy support (Nginx)

---

# Project Structure

```
go-weather-cli
│
├─ cmd
│   ├─ weather-cli
│   │   └─ main.go
│   │
│   └─ weather-server
│       └─ main.go
│
├─ internal
│   └─ weather
│       ├─ client.go
│       ├─ geo.go
│       ├─ weather.go
│       └─ types.go
│
├─ templates
│   └─ index.html
│
├─ go.mod
└─ README.md
```

---

# CLI Usage

Run the CLI version:

```
go run ./cmd/weather-cli Berlin
```

Example output:

```
City: Berlin
Temperature: 18°C
Feels like: 17°C
Humidity: 70%
Pressure: 1015 hPa
```

---

# Web Server

Run the web server:

```
go run ./cmd/weather-server
```

Open in browser:

```
http://localhost:8080
```

Enter a city name to see the current weather.


# Environment Variables

The application requires an OpenWeather API key.

```
export API_KEY=YOUR_API_KEY
```

---

# Deployment

Example deployment on a VPS:

```
git clone https://github.com/YOUR_USERNAME/go-weather-cli.git
cd go-weather-cli
go build -o weather-server ./cmd/weather-server
./weather-server
```

The application can also be run behind **Nginx reverse proxy**.

---

# Technologies

* Go
* OpenWeather API
* HTML templates
* Prometheus metrics
* Nginx
* systemd

---

# Future Improvements

* autocomplete for city search
* weather forecast
* Docker deployment
* Grafana dashboard
* caching
