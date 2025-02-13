package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/rltran-codex/myweather-api/internal/api/rest/models"
	"github.com/rltran-codex/myweather-api/internal/cache"
)

var (
	baseUrl string
	apiKey  string
)

func InitWeatherAPI() {
	baseUrl = "https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?key=%s"
	apiKey = os.Getenv("visual_crossing_api_key")
	if len(apiKey) == 0 {
		log.Fatal("ERROR - unable to parse environment variable 'visual_crossing_api_key'")
	}
}

func GetWeatherByCity(w http.ResponseWriter, r *http.Request) {
	// parse parameters
	city := r.URL.Query().Get("city")
	if len(city) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("no city was provided"))
		return
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		ip = r.RemoteAddr
	}
	log.Printf("'%s' requested weather data for city: %s", ip, city)

	// check cache
	if cache, err := checkCache(city); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(cache)
		return
	}

	// cache miss, run call API
	log.Println("cache missed, querying 'https://weather.visualcrossing.com/'")
	data := models.WeatherAPIResult{}
	query := fmt.Sprintf(baseUrl, city, apiKey)
	resp, err := http.Get(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("ERROR - %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("visualcrossing weather API returned Status: %d | Header: %v", resp.StatusCode, resp.Header)
	switch resp.StatusCode {
	case 200:
		// decode response body into data struct
		if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("ERROR - could not read response body and store into WeatherAPIResult struct: %v", err)
			return
		}

		// cache result
		if cacheData, err := json.Marshal(data); err == nil {
			cache.RClient.StoreCache(city, string(cacheData))
		} else {
			log.Printf("ERROR - cache failed: %v", err)
		}

		// send response back
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
		return
	default:
		body, _ := io.ReadAll(resp.Body)
		log.Printf("ERROR - [STATUS: %d] %s.", resp.StatusCode, string(body))
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("unable to fetch weather data, but its always sunny in Philadelphia"))
}

func checkCache(key string) (models.WeatherAPIResult, error) {
	data := models.WeatherAPIResult{}

	if cache, err := cache.RClient.CheckCache(key); err == nil {
		log.Printf("cache hit for %s", key)
		if err := json.Unmarshal([]byte(cache), &data); err != nil {
			fmt.Printf("ERROR - unable to unmarshall json: %v", err)
			return data, err
		}

		return data, nil
	}

	return data, fmt.Errorf("cache miss for %s", key)
}
