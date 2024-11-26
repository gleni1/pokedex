package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		Cache:   pokecache.NewCache(5 * time.Second),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := parts[0]
		args := parts[1:]
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
		if command == "explore" {
			if len(args) == 0 {
				fmt.Println("Usage: explore <area_name>")
				continue
			}
			areaName := args[0]
			pokeapi.CommandExplore(config, areaName)
		}
		if command == "catch" {
			if len(args) == 0 {
				fmt.Println("Usage: catch <pokemon_name>")
				continue
			}
			pokemonName := strings.Join(args, "")
			pokeapi.HandleCatch(config, pokemonName)
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
