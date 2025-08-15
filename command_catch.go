package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/HemahWeb/pokedexcli/internal/pokeapi"
)

func commandCatch(commandParam ...string) error {
	client := pokeapi.NewClient()
	pokemon, err := client.GetPokemon(commandParam[0])
	if err != nil {
		return errors.New("pokemon not found")
	}
	randomNumber := rand.Intn(100)
	captureRate := pokemon.BaseExperience / 2
	fmt.Printf("Throwing a Pokeball at %s...\n(%d%% chance of success)\n", pokemon.Name, captureRate)
	if randomNumber < captureRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		if _, ok := caughtPokemons[pokemon.Name]; !ok {
			fmt.Printf("Pokedex entry for %s was added. Type 'inspect %s' to see more details.\n", pokemon.Name, pokemon.Name)
		}
		caughtPokemons[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
