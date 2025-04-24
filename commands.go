package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/yashranjan1/pokedex/internal/pokecache"
)

type cli struct {
	config *config
	cache  *pokecache.Cache
}

func newCLI(conf *config, cache *pokecache.Cache) cli {
	return cli{config: conf, cache: cache}
}

const baseUrl string = "https://pokeapi.co/api/v2"

func (c cli) commandExit(inputs []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (c cli) commandHelp(inputs []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range c.getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func (c cli) commandMap(inputs []string) error {
	url := baseUrl + "/location-area"
	if c.config.next != "" {
		url = c.config.next
	}
	return helperMap(url, c.config, c.cache)
}

func (c cli) commandMapB(inputs []string) error {
	url := baseUrl + "/location-area"
	if c.config.prev != "" {
		url = c.config.prev
	}
	return helperMap(url, c.config, c.cache)
}

func helperMap(url string, config *config, cache *pokecache.Cache) error {

	data, exists := cache.Get(url)

	var body []byte
	if exists {
		body = data
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		cache.Add(url, body)
	}

	var result LocationRes
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result.Next != nil {
		config.next = *result.Next
	}
	if result.Previous != nil {
		config.prev = *result.Previous
	}
	for _, area := range result.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func (c cli) commandExplore(inputs []string) error {
	if len(inputs) != 1 {
		fmt.Println("Error: Incorrect usage of the explore command")
		fmt.Println("Usage:")
		fmt.Println("explore <area-name>")
		return nil
	}
	url := baseUrl + "/location-area/" + inputs[0]

	var body []byte
	if cachedData, exists := c.cache.Get(url); exists {
		body = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		c.cache.Add(url, body)
	}
	var result AreaRes
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Printf("Exploring %s...\n", inputs[0])
	fmt.Println("Found pokemon:")
	for _, pokemon := range result.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func (c cli) commandCatch(inputs []string) error {
	if len(inputs) != 1 {
		fmt.Println("Error: Incorrect usage of the explore command")
		fmt.Println("Usage:")
		fmt.Println("catch <area-name>")
		return nil
	}
	pokemon := inputs[0]
	url := baseUrl + "/pokemon/" + pokemon

	var body []byte
	if cachedData, exists := c.cache.Get(url); exists {
		body = cachedData
	} else {
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
			return err
		}

		defer res.Body.Close()

		if res.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
		}
		body, err = io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return err
		}
		c.cache.Add(url, body)
	}

	var result Pokemon
	err := json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return err
	}

	baseXP := float64(result.BaseExperience)
	chance := .75 * (1 - (.9 * (baseXP / 600)))

	randomNumber := rand.Float64()

	if randomNumber <= chance {
		fmt.Printf("%s was caught! \n", pokemon)
	} else {
		fmt.Printf("%s escaped! \n", pokemon)
	}
	return nil
}

func (c cli) getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    c.commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    c.commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display a list of the next 20 location area in the Pokemon world",
			callback:    c.commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display a list of the last 20 location area in the Pokemon world",
			callback:    c.commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Displays a list of all the pokemon in the area",
			callback:    c.commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    c.commandCatch,
		},
	}
}
