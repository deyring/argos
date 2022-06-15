package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Version      string        `yaml:"version"`
	Name         string        `yaml:"name"`
	Transactions []Transaction `yaml:"transactions"`
}

func (c *Config) Load(filename string) error {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(configFile, c); err != nil {
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

type Transaction struct {
	Name   string          `yaml:"name"`
	Checks []EndpointCheck `yaml:"checks"`
}
