package models

import (
	"errors"
	"strconv"
	"time"
)

type EndpointCheck struct {
	Name       string            `yaml:"name"`
	URL        string            `yaml:"url"`
	Method     Method            `yaml:"method"`
	Headers    map[string]string `yaml:"headers"`
	Body       string            `yaml:"body"`
	Timeout    int               `yaml:"timeout"`
	Assertions []Assertion       `yaml:"assertions"`
}

func (c *EndpointCheck) AssertResult(result *EndpointCheckResult) (success bool, err error) {
	success = true
	for _, assertion := range c.Assertions {
		switch assertion.Type {
		case AssertionTypeStatusCode:
			if assertion.Value != strconv.Itoa(result.StatusCode) {
				success = false
			}
		case AssertionTypeBody:
			if assertion.Value != string(result.Body) {
				success = false
			}
		case AssertionTypeHeader:
			err = errors.New("Header Assertion not implemented")
			return
		}
	}
	return
}

type Method string

const (
	MethodGet     Method = "GET"
	MethodHead    Method = "HEAD"
	MethodPost    Method = "POST"
	MethodPut     Method = "PUT"
	MethodPatch   Method = "PATCH"
	MethodDelete  Method = "DELETE"
	MethodConnect Method = "CONNECT"
	MethodOptions Method = "OPTIONS"
	MethodTrace   Method = "TRACE"
)

type EndpointCheckResult struct {
	StatusCode           int
	Body                 []byte
	Headers              map[string][]string
	DNSDuration          time.Duration
	TLSHandshakeDuration time.Duration
	ConnectDuration      time.Duration
	FirstByteDuration    time.Duration
	TotalDuration        time.Duration

	// Will only be set after successfull assertion
	Success bool
}
