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
