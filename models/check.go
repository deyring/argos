package models

import (
	"net/url"
	"strings"
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

func (c *EndpointCheck) Validate() error {
	if len(c.Name) == 0 {
		return ErrorCheckNameMissing
	}

	if len(c.URL) == 0 {
		return ErrorCheckURLMissing
	}

	u, err := url.ParseRequestURI(c.URL)
	if err != nil {
		return ErrorCheckURLInvalid
	}
	c.URL = u.String()

	c.Method = Method(strings.ToUpper(string(c.Method)))

	if c.Method != MethodGet &&
		c.Method != MethodHead &&
		c.Method != MethodPost &&
		c.Method != MethodPut &&
		c.Method != MethodPatch &&
		c.Method != MethodDelete &&
		c.Method != MethodConnect &&
		c.Method != MethodOptions &&
		c.Method != MethodTrace {
		return ErrorCheckInvalidMethod
	}

	if c.Assertions == nil || len(c.Assertions) == 0 {
		return ErrorCheckAssertionsMissing
	}

	return nil
}

func (c *EndpointCheck) AssertResult(result *EndpointCheckResult) (success bool, err error) {
	success = true
	for _, assertion := range c.Assertions {
		assertSuccess, err := assertion.AssertResult(result)
		if err != nil {
			return false, err
		}
		success = success && assertSuccess
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
	Name                 string
	StatusCode           int
	URL                  string
	Error                string
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
