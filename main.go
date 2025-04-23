package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yashranjan1/pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	config := &config{
		next: "",
		prev: "",
	}
	cache := pokecache.NewCache(5 * time.Second)
	c := newCLI(config, cache)
	commands := c.getCommands()
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		parts := cleanInput(text)

		opt, exists := commands[parts[0]]
		if !exists {
			fmt.Println("Unknown command")
		} else {
			opt.callback()
		}
	}
}

func cleanInput(text string) []string {
	lowered := strings.ToLower(text)
	split := strings.Fields(lowered)
	return split
}
