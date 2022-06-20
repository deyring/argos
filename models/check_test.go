package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckValidation(t *testing.T) {
	tests := []struct {
		description string
		check       EndpointCheck
		expected    error
	}{
		{
			description: "validates check",
			check: EndpointCheck{
				Name:    "test",
				URL:     "http://example.com",
				Method:  MethodGet,
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: nil,
		},
		{
			description: "failes on missing name",
			check: EndpointCheck{
				Name:    "",
				URL:     "http://example.com",
				Method:  MethodGet,
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckNameMissing,
		},
		{
			description: "failes on missing url",
			check: EndpointCheck{
				Name:    "test",
				URL:     "",
				Method:  MethodGet,
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckURLMissing,
		},
		{
			description: "failes on missing url scheme",
			check: EndpointCheck{
				Name:    "test",
				URL:     "www.example.com",
				Method:  MethodGet,
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckURLInvalid,
		},
		{
			description: "failes on invalid url",
			check: EndpointCheck{
				Name:    "test",
				URL:     "hi/there",
				Method:  MethodGet,
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckURLInvalid,
		},
		{
			description: "failes on invalid method",
			check: EndpointCheck{
				Name:    "test",
				URL:     "http://example.com",
				Method:  "invalid",
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckInvalidMethod,
		},
		{
			description: "failes on missing method",
			check: EndpointCheck{
				Name:    "test",
				URL:     "http://example.com",
				Method:  "",
				Headers: map[string]string{},
				Body:    "",
				Timeout: 0,
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			expected: ErrorCheckInvalidMethod,
		},
		{
			description: "failes on missing assertions",
			check: EndpointCheck{
				Name:       "test",
				URL:        "http://example.com",
				Method:     MethodGet,
				Headers:    map[string]string{},
				Body:       "",
				Timeout:    0,
				Assertions: []Assertion{},
			},
			expected: ErrorCheckAssertionsMissing,
		},
		{
			description: "failes on nil reference to assertions",
			check: EndpointCheck{
				Name:       "test",
				URL:        "http://example.com",
				Method:     MethodGet,
				Headers:    map[string]string{},
				Body:       "",
				Timeout:    0,
				Assertions: nil,
			},
			expected: ErrorCheckAssertionsMissing,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.check.Validate()
			if test.expected == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, test.expected.Error(), err.Error())
			}
		})
	}
}
