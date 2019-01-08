package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Version string
	Db      struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
	}
}

func ConfigInit() *Config {
	data, err := ioutil.ReadFile("env.yml")
	if err != nil {
		log.Fatal(err)
	}

	m := Config{}

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return &m
}
