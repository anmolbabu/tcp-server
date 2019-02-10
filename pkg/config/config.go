package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Server struct {
	Type     string `json:"connection_type"`
	Hostname string `json:"hostname"`
	Port     int    `json:"port"`
}

type Config struct {
	DataFile         string `json:"data_file"`
	Listener         Server `json:"listener"`
	Fetcher          Server `json:"fetcher"`
	FileDumpInterval int    `json:"async_file_dump_interval"`
}

func LoadConfiguration(file string) Config {
	var conf Config
	configData, err := ioutil.ReadFile(file)
	if err != nil {
		panic(fmt.Errorf("error parsing config file. Error: %v", err))
	}
	err = json.Unmarshal(configData, &conf)
	if err != nil {
		panic(fmt.Errorf("error parsing config file. Error: %v", err))
	}
	return conf
}
