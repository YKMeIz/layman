package main

import (
	"github.com/YKMeIz/layman/internal/pkgman"
	"log"
	"os"
)

func main() {

	args := os.Args
	var err error
	switch args[1] {
	case "-a":
		err = pkgman.Install(args[2:]...)
	case "-c":
		err = pkgman.Remove(args[2:]...)
	case "-u":
		pkgman.Update()
	case "-l":
		pkgman.List()
	default:
		os.Exit(1)
	}
	if err != nil {
		log.Println(err)
	}
}
