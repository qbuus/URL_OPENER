package main

import (
	"URL_OPENER/logger"
	"URL_OPENER/procedure"
	"os"

	"github.com/charmbracelet/huh"
)

func main() {
	for {
		logger.Info("<<<  URL OPENER  >>>")

		var OptionToSelect string

		form := huh.NewSelect[string]().
			Title("What would you like to do ?").Options(
			huh.NewOption("Test input string", "test_input_string"),
			huh.NewOption("Exit", "exit"),
		).Value(&OptionToSelect)

		err := form.Run()
		if err != nil {
			logger.Error("Error with selection", err.Error())
			continue
		}

		switch OptionToSelect {
		case "test_input_string":
			procedure.TestString()
		case "exit":
			logger.Info("Exiting...")
			os.Exit(0)
		default:
			logger.Error("Invalid option selected")
		}
	}
}
