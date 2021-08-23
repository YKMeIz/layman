package pkgman

import (
	"bufio"
	"fmt"
	"github.com/YKMeIz/layman/internal/color"
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
		fmt.Printf("%s [%s/%s]: ", color.Bold(s), color.Green("Yes"), color.Red("No"))

		response, err := reader.ReadString('\n')
		if err != nil {
			println(color.Red(err.Error()))
			os.Exit(1)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		switch response {
		case "y", "yes", "Y", "Yes":
			return true
		case "n", "no", "N", "No":
			return false
		default:
			continue
		}
	}
}
