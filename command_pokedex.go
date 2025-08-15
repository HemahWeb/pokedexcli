package main

import (
	"errors"
	"fmt"
	"sort"
)

func commandPokedex(commandParam ...string) error {
	if len(caughtPokemons) == 0 {
		return errors.New("no pokemons caught yet")
	}
	fmt.Println("Your Pokedex:")
	names := make([]string, 0, len(caughtPokemons))
	for name := range caughtPokemons {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
