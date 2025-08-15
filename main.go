package main

import "github.com/HemahWeb/pokedexcli/internal/pokeapi"

var caughtPokemons = make(map[string]pokeapi.PokemonFullResponse)

func main() {
	startRepl()
}
