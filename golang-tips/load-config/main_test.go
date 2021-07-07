package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	type Config struct {
		Test     string
		TestBool bool `envconfig:"TEST_BOOL"`
	}

	os.Setenv("TEST", "test")
	os.Setenv("TEST_BOOL", "true")

	config := Config{}
	err := LoadConfig("", &config)
	assert.NoError(t, err)

	assert.Equal(t, config.Test, "test")
	assert.Equal(t, config.TestBool, true)
}
