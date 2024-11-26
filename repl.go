package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/gleni1/pokedex/internal/pokeapi"
	"github.com/gleni1/pokedex/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &pokeapi.Config{
		Cache: pokecache.NewCache(5 * time.Second),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			commandExit()
		}
		if input == "help" {
			commands := getCommands()
			fmt.Println("Welcome to Pokedex!")
			fmt.Println("Usage: ")
			fmt.Println("help: " + commands["help"].description)
			fmt.Println("exit: " + commands["exit"].description)
			continue
		}
		if input == "map" {
			pokeapi.HandleMap(config)
			continue
		}
		if input == "mapb" {
			pokeapi.HandleBMap(config)
			continue
		}

		fmt.Println(input)
	}
}

func commandHelp() error {
	fmt.Println("here is the help function")
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
