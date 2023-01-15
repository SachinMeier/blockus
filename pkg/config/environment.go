package config

import (
	"os"

	"strconv"
)

type EnvProvider struct{}

func NewEnvProvider() Provider {
	return &EnvProvider{}
}

// GetInt64 returns an int64 value of the env var specified by keyName.
// If the env var is unset or an empty string or GetInt64 fails to
// parse an integer from the env var, the defaultVal is returned
func (e *EnvProvider) GetInt64(keyName string, defaultVal int64) int64 {
	val := os.Getenv(keyName)
	if val == "" {
		return defaultVal
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return defaultVal
	}
	return i
}

// GetString returns the value of the env var specified by keyName.
// If keyName is unset or an empty string, returns defaultVal.
func (e *EnvProvider) GetString(keyName string, defaultVal string) string {
	val := os.Getenv(keyName)
	if val == "" {
		return defaultVal
	}
	return val
}

// GetBool returns a boolean value of the env var specified by keyName.
// GetBool returns true if the env var is set to "true", "1", or the
// defaultVal is true. Otherwise, GetBool returns false.
func (e *EnvProvider) GetBool(keyName string, defaultVal bool) bool {
	val := os.Getenv(keyName)
	if val == "" {
		return defaultVal
	}
	return val == "true" || val == "1"
}
