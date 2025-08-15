package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/HemahWeb/pokedexcli/internal/pokecache"
)

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
	Cache      *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		BaseURL: "https://pokeapi.co/api/v2",
		Cache:   pokecache.NewCache(5 * time.Minute),
	}
}

func (c *Client) GetJSON(urlOrEndpoint string, result any) error {
	var url string
	if strings.HasPrefix(urlOrEndpoint, "http") {
		url = urlOrEndpoint // Full URL
	} else {
		url = c.BaseURL + urlOrEndpoint // Endpoint
	}

	if c.Cache != nil {
		if data, ok := c.Cache.Get(url); ok {
			return json.Unmarshal(data, result)
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return errors.New("failed to create request")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return errors.New("failed to get response")
	}
	defer resp.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = io.ReadAll(resp.Body)
	if err != nil {
		return errors.New("failed to read response body")
	}
	if c.Cache != nil && resp.StatusCode == 200 {
		c.Cache.Add(url, bodyBytes)
	}
	return json.Unmarshal(bodyBytes, result)
}

func (c *Client) ListLocations(pageURL *string) (LocationAreas, error) {
	urlOrEndpoint := "/location-area"
	if pageURL != nil {
		urlOrEndpoint = *pageURL
	}

	var locations LocationAreas
	err := c.GetJSON(urlOrEndpoint, &locations)
	if err != nil {
		return LocationAreas{}, errors.New("failed to get locations")
	}

	return locations, nil
}

func (c *Client) GetPokemonsAtLocation(locationName string) ([]Pokemon, error) {
	endpoint := "/location-area/" + locationName
	var response LocationAreasFullResponse
	err := c.GetJSON(endpoint, &response)
	if err != nil {
		return []Pokemon{}, err
	}
	var pokemons []Pokemon
	for _, encounter := range response.PokemonEncounters {
		pokemons = append(pokemons, Pokemon{
			Name: encounter.Pokemon.Name,
			URL:  encounter.Pokemon.URL,
		})
	}
	return pokemons, nil
}

func (c *Client) GetPokemon(name string) (PokemonFullResponse, error) {
	endpoint := "/pokemon/" + name
	var response PokemonFullResponse
	err := c.GetJSON(endpoint, &response)
	if err != nil {
		return PokemonFullResponse{}, err
	}
	return response, nil
}
