package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var englishRights string

//go:embed french_rights.txt
var frenchRights string

var languages = map[string]bool{
	"english": true,
	"french":  true,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: Missing command-line argument")
		os.Exit(1)
	}

	arg := os.Args[1]
	fmt.Println("Argument provided:", arg)

	if _, ok := languages[arg]; !ok {
		fmt.Println("Error: Invalid language")
		os.Exit(1)
	}

	switch arg {
	case "english":
		fmt.Println(englishRights)
	case "french":
		fmt.Println(frenchRights)
	default:
		fmt.Println("Error: Invalid language", arg)
	}

}
