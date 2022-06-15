package sinkfactory

import (
	resultsink "github.com/deyring/argos/resultsink"
	"github.com/deyring/argos/resultsink/sinks/stdout"
)

func GetNewStdoutSink() resultsink.Sink {
	return stdout.New()
}
