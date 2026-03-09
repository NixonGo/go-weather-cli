package weather

import (
	"net/http"
	"os"
	"time"
)

var apiKey = os.Getenv("API_KEY")

var httpClient = &http.Client{Timeout: 10 * time.Second}
