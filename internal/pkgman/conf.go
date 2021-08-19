package pkgman

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	askMode, verboseMode, skippgpcheck bool
)

func SetAskMode() {
	askMode = true
}

func SetVerboseMode() {
	verboseMode = true
}

func SetSkipPGPCheck() {
	skippgpcheck = true
}

func askForConfirmation(s string) bool {
	if !askMode {
		return true
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("%s [Yes/No]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" || response == "Y" || response == "Yes" {
			return true
		} else if response == "n" || response == "no" || response == "N" || response == "No" {
			return false
		}
	}
}
