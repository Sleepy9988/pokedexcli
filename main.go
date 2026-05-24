package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}

	for {
		fmt.Print("Pokedex > ")

		for scanner.Scan() {
			user_input := scanner.Text()
			cleaned_input := cleanInput(user_input)
			fmt.Printf("Your command was: %s\n", cleaned_input[0])
			break
		}
	}
}
