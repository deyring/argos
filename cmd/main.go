package main

import (
	"io"
	"os"

	"github.com/deyring/argos/runner"
	"github.com/deyring/argos/utils"
)

func main() {
	logger := utils.NewLogrusLogger()
	logger.Info("starting argos...")

	configFileReader, err := openStdinOrFile()
	if err != nil {
		logger.Fatalf("failed to open config file: %v", err)
	}

	runner, err := runner.New(logger, configFileReader)
	if err != nil {
		logger.Fatalf("failed to create runner: %v", err)
	}

	if err := runner.Run(); err != nil {
		logger.Fatalf("failed to run runner: %v", err)
	}
}

func openStdinOrFile() (io.Reader, error) {
	var err error
	r := os.Stdin
	if len(os.Args) > 1 {
		r, err = os.Open(os.Args[1])
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}
