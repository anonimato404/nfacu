package cmd

import (
	"flag"
	"os"
)

func getArguments() (bool, string) {
	help := flag.Bool("help", false, "Show help")

	flag.Parse()

	if len(os.Args) > 1 {
		return *help, os.Args[1]
	}

	return *help, "nfacu.json"
}
