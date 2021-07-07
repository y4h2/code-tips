package main

import (
	"github.com/pkg/errors"

	"github.com/kelseyhightower/envconfig"
)

func LoadConfig(prefix string, v interface{}) error {
	err := envconfig.Process(prefix, v)
	if err != nil {
		return errors.Wrap(err, "failed to load config")
	}

	return nil
}
