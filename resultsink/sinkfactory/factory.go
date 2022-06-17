package sinkfactory

import (
	resultsink "github.com/deyring/argos/resultsink"
	"github.com/deyring/argos/resultsink/sinks/influxdb"
	"github.com/deyring/argos/resultsink/sinks/stdout"
)

func GetNewStdoutSink() resultsink.Sink {
	return stdout.New()
}

func GetNewInfluxDBSink(host, user, password, databse string) resultsink.Sink {
	return influxdb.New(host, user, password, databse)
}
