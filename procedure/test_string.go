package procedure

import (
	"URL_OPENER/logger"

	"github.com/charmbracelet/huh"
)

func TestString() {
	logger.Info("Test input string")

	var testString string
	title := "Test input string"
	description := "Leave empty to return to menu"

	err := huh.NewInput().
		Title(title).
		Description(description).
		Prompt("?").
		Value(&testString).
		Run()

	if err != nil {
		logger.Error("Error reading input:", err)
		return
	}

	if testString == "" {
		logger.Info("Returning to menu...")
		return
	}

	logger.Info("You entered:", testString)
}
