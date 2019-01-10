package config

import (
	"io/ioutil"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Version string
	Db      struct {
		Host     string
		Port     uint16
		Database string
		Username string
		Password string
	}
	Time string
}

var c *config
var once sync.Once

func init() {
	once.Do(func() {
		data, _ := ioutil.ReadFile("env.yml")

		c := config{}
		_ = yaml.Unmarshal([]byte(data), &c)
		Instance = &c
	})
}

// func ConfigInit() *config {

// 	once.Do(func() {
// 		data, _ := ioutil.ReadFile("env.yml")

// 		c := config{}
// 		_ = yaml.Unmarshal([]byte(data), &c)
// 		Instance = &c
// 	})

// 	return Instance
// }

var Instance *config
