package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"go-project/internal/weather"
)

func main() {
	w, _ := RequestWeatherAPI()

	if w.Weather[0].IsUmbrellaNeeded() {
		fmt.Println("傘持って外出したかい？")
	} else {
		fmt.Println("今日は傘必要ないかもね！")
	}
}

func RequestWeatherAPI() (weather.CurrentWeatherStatus, error) {
	reqURL := generateRequest()

	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var w weather.CurrentWeatherStatus
	if err := json.NewDecoder(resp.Body).Decode(&w); err != nil {
		log.Fatal(err)
	}

	return w, err
}

func generateRequest() string {
	apiURL := "https://api.openweathermap.org/data/2.5/weather"
	params := map[string]string{
		"lat":   "35.6983223",
		"lon":   "139.7730186",
		"appid": "0146307c9953096ab5b19594779feac0",
	}
	queryParams := url.Values{}
	for k, v := range params {
		queryParams.Set(k, v)
	}

	return fmt.Sprintf("%s?%s", apiURL, queryParams.Encode())
}
