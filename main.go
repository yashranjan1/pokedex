package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	return strings.Fields(text)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		parts := strings.Fields(strings.ToLower(text))
		fmt.Printf("Your command was: %s\n", parts[0])
	}
}
