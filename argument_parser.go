package main

import (
	"flag"
	"strings"
	"unicode"
)

type Arguments struct {
	UPSManifestPath string
	UPSToDeploy []string
}

func ParseArguments(args []string) Arguments {
	arguments := Arguments{}

	f := flag.NewFlagSet("deploy-ups", flag.ExitOnError)

	f.StringVar(&arguments.UPSManifestPath, "f", "", "")
	var upses string
	f.StringVar(&upses, "u", "", "")
	f.Parse(args[1:])

	if upses != "" {
		arguments.UPSToDeploy = strings.Split(strings.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return -1
			} else if r == ';' {
				return ','
			}
			return r
		}, upses), ",")
	}
	return arguments
}