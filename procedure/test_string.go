package procedure

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func TestString() {
	fmt.Println("test string")

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
		fmt.Println("Error reading input:", err)
		return
	}

	if testString == "" {
		fmt.Println("No input. Returning to menu...")
		return
	}

	fmt.Println("You entered:", testString)
}
