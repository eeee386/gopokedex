package commands

import (
	"GoPokedex/service"
	"errors"
	"fmt"
	"math/rand"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(params ...string) error
}

var cliMap = make(map[string]CliCommand)

func commandHelp(params ...string) error {
	if cliMap == nil {
		return errors.New("No commands to show")
	}
	fmt.Println("Welcome to Pokedex!")
	fmt.Print("Usage:\n\n")
	for k, v := range cliMap {
		fmt.Printf("%s: %s\n", k, v.Description)
	}
	fmt.Print("\n")
	return nil
}

func commandExit(params ...string) error {
	fmt.Println("Exiting pokedex")
	return nil
}


var prevLocationsUrl string
var nextLocationUrl = "https://pokeapi.co/api/v2/location-area"

func commandMap(params ...string) error {
	if(nextLocationUrl == "") {
		fmt.Println("No more available locations to show")
		return nil
	}
  locationObject := service.GetLocation(nextLocationUrl)
	for _,l := range locationObject.Results {
		fmt.Println(l.Name)
	}
	prevLocationsUrl = locationObject.Previous
	nextLocationUrl = locationObject.Next
	return nil
}

func commandMapB (params ...string) error {
	if(prevLocationsUrl == "") {
		return errors.New("There no previous location records")
	}
	var locationsObject = service.GetLocation(prevLocationsUrl)
	for _,l := range locationsObject.Results {
		fmt.Println(l.Name)
	}
	prevLocationsUrl = locationsObject.Previous
	nextLocationUrl = locationsObject.Next
	return nil
}

func commandExplore(params ...string) error {
	var locationAreaObject = service.GetLocationArea(params[0])
	var pokemons = locationAreaObject.PokemonEncounters
	for _, v := range(pokemons) {
		fmt.Println(v.Pokemon.Name)
	}
	return nil
}

var pokedex = make(map[string]service.PokemonResponseObject)

func commandCatch(params ...string) error {
	var pokemonObject = service.GetPokemon(params[0])
	pokemon, ok := pokedex[pokemonObject.Name]
	if(ok) {
		fmt.Printf("%s is already in the pokedex\n", pokemon.Name)
		return nil
	}
	fmt.Printf("Throwing a pokeball at %s...\n", pokemonObject.Name)
	catchNumber := rand.Intn(1000)
	if catchNumber >= pokemonObject.BaseExperience {
		pokedex[pokemonObject.Name] = pokemonObject
		fmt.Printf("%s was caught!\n", pokemonObject.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemonObject.Name)
	}
	return nil
	
}

func GetCLICommands() map[string]CliCommand {
	cliMap = map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Display help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name: "map",
			Description: "Show (the next) 20 locations",
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Show the previous 20 locations",
			Callback: commandMapB,
		},
		"explore": {
			Name: "explore",
			Description: "Show pokemons in an area",
			Callback: commandExplore,
		},
		"catch": {
			Name: "catch",
			Description: "Catch a pokemon to your pokedex",
			Callback: commandCatch,
		},
	}
	return cliMap
}
