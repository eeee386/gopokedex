package commands

import (
	"fmt"
	"errors"
	"GoPokedex/service"
	"GoPokedex/pokecache"
	"time"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var cache = pokecache.NewCache(300 * time.Second)

var cliMap = make(map[string]CliCommand)

func commandHelp() error {
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

func commandExit() error {
	fmt.Println("Exiting pokedex")
	return nil
}


var prevLocationsUrl string
var nextLocationUrl = "https://pokeapi.co/api/v2/location-area"

func commandMap() error {
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

func commandMapB () error {
	fmt.Println("prev: ", prevLocationsUrl == "")
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
	}
	return cliMap
}
