package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// require.Equal(t, 8080, conf.Server.Port, "Server.Port should be 8080")
//
// TestLoadConfig test AppConfig with ./config.yaml
func TestLoadConfig(t *testing.T) {
	configFile := "./config.yaml"

	_, err := LoadConfig[AppConfig](configFile)

	require.NoErrorf(t, err, "Failed to load config: %v", err)
}
