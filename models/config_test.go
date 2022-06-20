package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {

	t.Run("LoadsConfig", func(t *testing.T) {
		config := &Config{}
		file, err := os.Open("../examples/simple.yml")
		require.NoError(t, err)
		err = config.Load(file)
		require.NoError(t, err)
	})

}

func TestTransaction(t *testing.T) {
	tests := []struct {
		description string
		transaction Transaction
		expected    error
	}{
		{
			description: "validates transaction",
			transaction: Transaction{
				Name: "test",
				Checks: []EndpointCheck{
					{
						Name: "test",
					},
				},
			},
			expected: nil,
		},
		{
			description: "failes on missing name",
			transaction: Transaction{
				Name: "",
				Checks: []EndpointCheck{
					{
						Name: "test",
					},
				},
			},
			expected: ErrorTransactionNameMissing,
		},
		{
			description: "failes on missing checks",
			transaction: Transaction{
				Name:   "test",
				Checks: []EndpointCheck{},
			},
			expected: ErrorTransactionChecksMissing,
		},
		{
			description: "failes on nil reference to checks",
			transaction: Transaction{
				Name:   "test",
				Checks: nil,
			},
			expected: ErrorTransactionChecksMissing,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			err := test.transaction.Validate()
			if test.expected == nil {
				require.NoError(t, err)
			} else {
				require.Equal(t, test.expected.Error(), err.Error())
			}
		})
	}
}
