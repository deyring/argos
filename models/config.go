package models

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version      string        `yaml:"version"`
	Name         string        `yaml:"name"`
	Execute      ExecuteType   `yaml:"execute"`
	Sleep        int           `yaml:"sleep"`
	Transactions []Transaction `yaml:"transactions"`
	Outputs      []Output      `yaml:"outputs"`
}

type ExecuteType string

const (
	ExecuteTypeOnce ExecuteType = "once"
	ExecuteTypeLoop ExecuteType = "loop"
)

func (c *Config) Load(configFileReader io.Reader) error {

	fileContent, err := ioutil.ReadAll(configFileReader)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(fileContent, c); err != nil {
		return err
	}

	return nil
}

func (c *Config) Save(filename string) error {
	configFile, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, configFile, 0644)
}

type Result struct {
	Name               string
	TransactionResults []TransactionResult
}

type Transaction struct {
	Name   string          `yaml:"name"`
	Checks []EndpointCheck `yaml:"checks"`
}

type TransactionResult struct {
	Name                 string
	Success              bool
	EndpointCheckResults []EndpointCheckResult
}

type OutputType string

const (
	OutputTypeStdOut   OutputType = "stdout"
	OutputTypeInfluxDB OutputType = "influxdb"
)

type Output struct {
	Type     OutputType `yaml:"type"`
	Host     string     `yaml:"host"`
	User     string     `yaml:"user"`
	Password string     `yaml:"password"`
	Database string     `yaml:"database"`
}
