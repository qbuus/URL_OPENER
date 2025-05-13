package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/qbuus/URL_OPENER/procedure"
)

func main() {
	for {
		var OptionToSelect string

		form := huh.NewSelect[string]().
			Title("What would you like to do ?").Options(
			huh.NewOption("Test input string", "test_input_string"),
			huh.NewOption("Exit", "exit"),
		).Value(&OptionToSelect)

		err := form.Run()
		if err != nil {
			fmt.Println("Error with selection", err.Error())
			continue
		}

		switch OptionToSelect {
		case "test_input_string":
			procedure.TestString()
		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid option selected")
		}
	}
}
