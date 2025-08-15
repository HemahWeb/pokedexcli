package main

import (
	"errors"
	"fmt"

	"github.com/HemahWeb/pokedexcli/internal/pokeapi"
)

var (
	mapPrevURL     *string
	mapNextURL     *string
	mapCurrentPage int = 1
	mapTotalPages  int = 1
)

func commandMapf(commandParam ...string) error {
	client := pokeapi.NewClient()
	locations, err := client.ListLocations(mapNextURL)
	if err != nil {
		return errors.New("failed to get next page")
	}

	mapPrevURL = locations.Previous
	mapNextURL = locations.Next

	// Calculate total pages (PokeAPI default page size is 20)
	if locations.Count > 0 {
		mapTotalPages = (locations.Count + 19) / 20
	}
	if mapPrevURL == nil {
		mapCurrentPage = 1
	} else if mapNextURL != nil && mapCurrentPage < mapTotalPages {
		mapCurrentPage++
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Printf("\nPage %d of %d\n", mapCurrentPage, mapTotalPages)
	return nil
}

func commandMapb(commandParam ...string) error {
	if mapPrevURL == nil {
		return errors.New("you are on the first page")
	}
	client := pokeapi.NewClient()
	locations, err := client.ListLocations(mapPrevURL)
	if err != nil {
		return errors.New("failed to get previous page")
	}

	mapPrevURL = locations.Previous
	mapNextURL = locations.Next

	if mapCurrentPage > 1 {
		mapCurrentPage--
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Printf("\nPage %d of %d\n", mapCurrentPage, mapTotalPages)
	return nil
}
