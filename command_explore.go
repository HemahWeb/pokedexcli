package main

import (
	"errors"
	"fmt"

	"github.com/HemahWeb/pokedexcli/internal/pokeapi"
)

func commandExplore(commandParam ...string) error {
	client := pokeapi.NewClient()
	pokemons, err := client.GetPokemonsAtLocation(commandParam[0])
	if err != nil {
		return errors.New("location not found")
	}
	fmt.Printf("Pokemons at %s:\n", commandParam)
	for _, pokemon := range pokemons {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
