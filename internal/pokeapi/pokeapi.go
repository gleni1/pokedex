package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gleni1/pokedex/internal/pokecache"
)

type Config struct {
	NextURL     *string
	PreviousURL *string
	Cache       *pokecache.Cache
}

type APIResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func HandleMap(config *Config) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.NextURL != nil {
		url = *(config.NextURL)
	}

	response, err := FetchLocationAreas(url, config)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	DisplayLocations(response.Results)

	config.NextURL = response.Next
	config.PreviousURL = response.Previous
}

func HandleBMap(config *Config) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if config.PreviousURL != nil {
		url = *config.PreviousURL
	}

	response, err := FetchLocationAreas(url, config)
	if err != nil {
		fmt.Println("Error fetching locations:", err)
		return
	}

	DisplayLocations(response.Results)

	config.NextURL = response.Next
	config.PreviousURL = response.Previous
}

func FetchLocationAreas(url string, config *Config) (*APIResponse, error) {
	cachedData, found := config.Cache.Get(url)
	if found {
		var apiResponse APIResponse
		err := json.Unmarshal(cachedData, &apiResponse)
		if err != nil {
			return nil, err
		}
		return &apiResponse, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse APIResponse

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}

func DisplayLocations(locations []LocationArea) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func HandleMapBack(config *Config) {

}
