package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type config struct {
	API struct {
		BaseURL   string `yaml:"baseURL"`
		Key       string `yaml:"key"`
		ShopID    int    `yaml:"shopID"`
		UserPrefs string `yaml:"userPrefs"`
	} `yaml:"api"`

	CloudSQLDev struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"database_name"`
	} `yaml:"cloud-sql-dev"`

	LocalSQLDev struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"database_name"`
	} `yaml:"local-sql-dev"`

	MongoLocal struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Address  string `yaml:"address"`
		DBName   string `yaml:"database_name"`
	} `yaml:"mongo-local"`
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
