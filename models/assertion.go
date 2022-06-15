package models

type AssertionType string

const (
	AssertionTypeStatusCode AssertionType = "status_code"
	AssertionTypeBody       AssertionType = "body"
	AssertionTypeHeader     AssertionType = "header"
)

type Assertion struct {
	Name  string        `yaml:"name"`
	Type  AssertionType `yaml:"type"`
	Value string        `yaml:"value"`
}
