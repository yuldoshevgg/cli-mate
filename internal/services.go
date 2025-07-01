package internal

import (
	"cli-mate/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type AirPollutionResponse struct {
	Coord Coord  `json:"coord"`
	List  []List `json:"list"`
}

type AirPollutionMain struct {
	Aqi int `json:"aqi"`
}

type Components struct {
	Co   float64 `json:"co"`
	No   float64 `json:"no"`
	No2  float64 `json:"no2"`
	O3   float64 `json:"o3"`
	So2  float64 `json:"so2"`
	Pm25 float64 `json:"pm2_5"`
	Pm10 float64 `json:"pm10"`
	Nh3  float64 `json:"nh3"`
}

type List struct {
	Main       AirPollutionMain `json:"main"`
	Components Components       `json:"components"`
	Dt         int              `json:"dt"`
}

type AirPollutionType struct {
	Color string `json:"color"`
	Emoji string `json:"emoji"`
	Text  string `json:"text"`
}

var AirPollutionMainColor = map[int]AirPollutionType{
	1: {Emoji: "🟢", Color: "\033[32m", Text: "Good"},
	2: {Emoji: "🟡", Color: "\033[33m", Text: "Normal"},
	3: {Emoji: "🟠", Color: "\u001b[38;2;253;182;0m", Text: "Moderate"},
	4: {Emoji: "🔴", Color: "\033[31m", Text: "Poor"},
	5: {Emoji: "🟣", Color: "\033[35m", Text: "Very Poor"},
}

func AirPollution(cfg *config.Config, lat float64, lon float64) (*AirPollutionResponse, error) {
	url := fmt.Sprintf("%v/data/2.5/air_pollution?lat=%v&lon=%v&appid=%v", cfg.URL, lat, lon, cfg.APIKey)
	data, err := MakeRequest(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	var response AirPollutionResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type WeatherResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int       `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func GetWeather(cfg *config.Config, lat float64, lon float64) (*WeatherResponse, error) {
	url := fmt.Sprintf("%v/data/2.5/weather?lat=%v&lon=%v&appid=%v&units=metric", cfg.URL, lat, lon, cfg.APIKey)
	data, err := MakeRequest(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	var response WeatherResponse
	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

type GeoLocationResponse struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	Query       string  `json:"query"`
}

func GetGeoLocation() (*GeoLocationResponse, error) {
	var (
		response *GeoLocationResponse
	)

	data, err := MakeRequest("http://ip-api.com/json", http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return response, nil
}
