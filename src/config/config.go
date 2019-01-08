package config

import (
	"io/ioutil"
	"log"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Version string
	Db      struct {
		Host     string
		Port     int
		Database string
		Username string
		Password string
	}
}

var c *config
var once sync.Once

func init() {
	data, err := ioutil.ReadFile("env.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	if c == nil {
		once.Do(func() {
			c := config{}

			err = yaml.Unmarshal([]byte(data), &c)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			Instance = &c
		})
	}
}

var Instance *config
