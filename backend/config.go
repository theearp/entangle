package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	SQL struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"database_name"`
	} `yaml:"sql"`
}

func getConfig(fileName string) (*config, error) {
	log.Println("fetching config..")
	var c *config
	conf, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read %q, aborting: %s", fileName, err)
	}
	err = yaml.Unmarshal(conf, &c)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %q: %s", fileName, err)
	}
	return c, nil
}
