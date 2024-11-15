package server

import (
	"cli-mate/server/config"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAirPollution() {}

func GetGeoLocation(cfg *config.Config, location string) ([]string, error) {
	var (
		response []string
	)

	url := fmt.Sprintf("%v/geo/1.0/direct?q=%v&appid=%v", cfg.WeatherURL, location, cfg.APIKey)
	data, err := MakeRequest(url, http.MethodGet, nil, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("data: ", data)

	if err = json.Unmarshal(data, &response); err != nil {
		return nil, err
	}

	return response, nil
}
