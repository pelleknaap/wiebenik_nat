package models

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"gopkg.in/yaml.v2"
	"os"
)

type Config struct {
	Address string `yaml:"addr"`
}

func (c *Config) Parse() error {
	var path string
	if os.Getenv("PROD") == "true" {
		path = "config_prod.yml"
	} else {
		path = "config_dev.yml"
	}

	f, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("error opening config file: %s", err))
	}

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(fmt.Sprintf("error decoding yaml from file: %s", err))
	}

	spew.Dump(c)

	return nil
}