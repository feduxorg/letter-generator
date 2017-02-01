package main

import (
	"github.com/maxmeyer/letter-generator-go/commandline"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr, could also be a file.
	// log.SetOutput(os.Stdout)
}

func main() {
	cli := commandline.Cli{}
	err := cli.Run(os.Args)

	if err != nil {
		os.Exit(1)
	}
}
