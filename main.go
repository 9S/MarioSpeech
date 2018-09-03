package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

// Variables used for command line parameters
var (
	text string
)

func init() {
	cliArgs := os.Args[1:]
	text = strings.Join(cliArgs, " ")
}

func main() {
	const SPACE string = ":nothing:"

	// Convert Text to Uppercase
	text = strings.ToUpper(text)

	// Replace everything not A-Z with nothing
	regex := regexp.MustCompile("[^A-Z ]+")
	processedString := regex.ReplaceAllString(text, "")

	// Replace every letter with it's big mario emote
	regex = regexp.MustCompile(`(\w)`)
	processedString = regex.ReplaceAllString(processedString, ":mar$1:")

	// Replace every space with the contents of SPACE
	processedString = strings.Replace(processedString, " ", SPACE, -1)

	fmt.Println(processedString)

	err := clipboard.WriteAll(processedString)
	if err != nil {
		fmt.Println("Could not copy to clipboard.")
	}
}
