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

type LocationAreaDetails struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
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

func CommandExplore(config *Config, areaName string) {
	fmt.Printf("Exploring %s...\n", areaName)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", areaName)

	response, err := FetchLocationAreaDetails(url, config)
	if err != nil {
		fmt.Println("Error fetching location area details:", err)
		return
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
}

func FetchLocationAreaDetails(url string, config *Config) (*LocationAreaDetails, error) {
	cachedData, found := config.Cache.Get(url)
	if found {
		var details LocationAreaDetails
		err := json.Unmarshal(cachedData, &details)
		if err != nil {
			return nil, err
		}
		return &details, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var details LocationAreaDetails
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&details)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(details)
	if err == nil {
		config.Cache.Add(url, data)
	}

	return &details, nil

}

func DisplayLocations(locations []LocationArea) {
	for _, location := range locations {
		fmt.Println(location.Name)
	}
}

func HandleMapBack(config *Config) {

}
