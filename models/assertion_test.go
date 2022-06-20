package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAssertion(t *testing.T) {
	tests := []struct {
		description   string
		assertion     Assertion
		expectSuccess bool
		expected      error
	}{
		{
			description: "positive status code assertion",
			assertion: Assertion{
				Type:  AssertionTypeStatusCode,
				Value: "200",
			},
			expectSuccess: true,
			expected:      nil,
		},
		{
			description: "negative status code assertion",
			assertion: Assertion{
				Type:  AssertionTypeStatusCode,
				Value: "404",
			},
			expectSuccess: false,
			expected:      nil,
		},
		{
			description: "positive body assertion",
			assertion: Assertion{
				Type:  AssertionTypeBody,
				Value: "test",
			},
			expectSuccess: true,
			expected:      nil,
		},
		{
			description: "negative body assertion",
			assertion: Assertion{
				Type:  AssertionTypeBody,
				Value: "test2",
			},
			expectSuccess: false,
			expected:      nil,
		},
		{
			description: "positive header assertion",
			assertion: Assertion{
				Type:  "AssertionTypeHeader",
				Value: "test",
			},
			expectSuccess: false,
			expected:      ErrorNotImplemented,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := &EndpointCheckResult{
				StatusCode: 200,
				Body:       []byte("test"),
			}
			success, err := test.assertion.AssertResult(result)
			require.Equal(t, test.expectSuccess, success)
			if test.expected != nil {
				require.Equal(t, test.expected.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAssertResult(t *testing.T) {
	tests := []struct {
		description   string
		check         *EndpointCheck
		result        *EndpointCheckResult
		expectSuccess bool
		expected      error
	}{
		{
			description: "positive status code assertion",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
				},
			},
			result: &EndpointCheckResult{
				StatusCode: 200,
			},
			expectSuccess: true,
			expected:      nil,
		},
		{
			description: "negative status code assertion",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "404",
					},
				},
			},
			result: &EndpointCheckResult{
				StatusCode: 200,
			},
			expectSuccess: false,
			expected:      nil,
		},
		{
			description: "positive body assertion",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeBody,
						Value: "test",
					},
				},
			},
			result: &EndpointCheckResult{
				Body: []byte("test"),
			},
			expectSuccess: true,
			expected:      nil,
		},
		{
			description: "negative body assertion",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeBody,
						Value: "test2",
					},
				},
			},
			result: &EndpointCheckResult{
				Body: []byte("test"),
			},
			expectSuccess: false,
			expected:      nil,
		},
		{
			description: "positive multiple assertion",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
					{
						Type:  AssertionTypeBody,
						Value: "test",
					},
				},
			},
			result: &EndpointCheckResult{
				StatusCode: 200,
				Body:       []byte("test"),
			},
			expectSuccess: true,
			expected:      nil,
		},
		{
			description: "negative multiple assertion, wrong body",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "200",
					},
					{
						Type:  AssertionTypeBody,
						Value: "test2",
					},
				},
			},
			result: &EndpointCheckResult{
				StatusCode: 200,
				Body:       []byte("test"),
			},
			expectSuccess: false,
			expected:      nil,
		},
		{
			description: "negative multiple assertion, wrong status code",
			check: &EndpointCheck{
				Assertions: []Assertion{
					{
						Type:  AssertionTypeStatusCode,
						Value: "404",
					},
					{
						Type:  AssertionTypeBody,
						Value: "test",
					},
				},
			},
			result: &EndpointCheckResult{
				StatusCode: 200,
				Body:       []byte("test"),
			},
			expectSuccess: false,
			expected:      nil,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			success, err := test.check.AssertResult(test.result)
			require.Equal(t, test.expectSuccess, success)
			if test.expected != nil {
				require.Equal(t, test.expected.Error(), err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
