package models

import (
	"strconv"
)

type AssertionType string

const (
	AssertionTypeStatusCode AssertionType = "status_code"
	AssertionTypeBody       AssertionType = "body"
)

type Assertion struct {
	Name  string        `yaml:"name"`
	Type  AssertionType `yaml:"type"`
	Value string        `yaml:"value"`
}

func (a *Assertion) AssertResult(result *EndpointCheckResult) (success bool, err error) {
	success = true
	switch a.Type {
	case AssertionTypeStatusCode:
		success, err = a.runForTypeStatusCode(result)
	case AssertionTypeBody:
		success, err = a.runForTypeBody(result)
	default:
		err = ErrorNotImplemented
		success = false
		return
	}
	return
}

func (a *Assertion) runForTypeStatusCode(result *EndpointCheckResult) (success bool, err error) {
	success = true
	if a.Value != strconv.Itoa(result.StatusCode) {
		success = false
	}
	return
}

func (a *Assertion) runForTypeBody(result *EndpointCheckResult) (success bool, err error) {
	success = true
	if a.Value != string(result.Body) {
		success = false
	}
	return
}
