package main

import (
	"flag"

	"github.com/deyring/argos/runner"
	"github.com/deyring/argos/utils"
)

func main() {
	logger := utils.NewLogrusLogger()
	logger.Info("starting argos...")

	flags := parseFlags()
	logger.Debugf("parsed flags: %v", flags)

	runner, err := runner.New(logger, flags.Filename)
	if err != nil {
		logger.Fatalf("failed to create runner: %v", err)
	}

	if err := runner.Run(); err != nil {
		logger.Fatalf("failed to run runner: %v", err)
	}
}

func parseFlags() *flags {
	filenamePtr := flag.String("f", "argos.yml", "filename of the argos config file")
	// Parse flags
	flag.Parse()

	return &flags{
		Filename: *filenamePtr,
	}
}
