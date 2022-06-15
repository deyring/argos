package models

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {

	t.Run("LoadsConfig", func(t *testing.T) {
		config := &Config{}
		err := config.Load("../examples/simple.yml")
		require.NoError(t, err)
	})

	t.Run("SavesConfig", func(t *testing.T) {
		config := &Config{
			Version: "1.0.0",
			Transactions: []Transaction{
				{
					Name: "Google Search",
					Checks: []EndpointCheck{
						{
							Name:    "index",
							URL:     "https://www.google.com/",
							Method:  MethodGet,
							Timeout: 5,
							Assertions: []Assertion{
								{
									Name:  "status",
									Type:  AssertionTypeStatusCode,
									Value: "200",
								},
							},
						},
					},
				},
			},
		}
		err := config.Save("../examples/configWrite.yml")
		require.NoError(t, err)
	})

	t.Run("SavesAndReadsConfig", func(t *testing.T) {
		filename := "../examples/configWrite.yml"
		config := &Config{
			Version: "1.0.0",
			Transactions: []Transaction{
				{
					Name: "Google Search",
					Checks: []EndpointCheck{
						{
							Name:    "index",
							URL:     "https://www.google.com/",
							Method:  MethodGet,
							Headers: map[string]string{"User-Agent": "argos test client"},
							Timeout: 5,
							Assertions: []Assertion{
								{
									Name:  "status",
									Type:  AssertionTypeStatusCode,
									Value: "200",
								},
							},
						},
					},
				},
			},
		}
		err := config.Save(filename)
		require.NoError(t, err)

		config2 := &Config{}
		err = config2.Load(filename)
		require.NoError(t, err)
		require.EqualValues(t, config, config2)
	})

}
