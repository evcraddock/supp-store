package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type WebsiteConfig struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type DbConfig struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}

type Configuration struct {
	Website WebsiteConfig `json:"website"`
	MongoDb DbConfig      `json:"mongodb"`
}

var configFile = flag.String("conf", "suppstore.conf", "load configuration file")

func LoadFile() *Configuration {
	flag.Parse()

	configuration, err := ioutil.ReadFile(*configFile)
	if err != nil {
		panic(err)
	}

	config := &Configuration{}
	err = json.Unmarshal(configuration, config)
	if err != nil {
		panic(err)
	}

	return config
}
