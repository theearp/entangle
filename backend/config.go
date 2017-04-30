package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v2"
)

type config struct {
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
}

func (c *config) env(env string) (*mysql.Config, error) {
	var err error
	var cfg *mysql.Config
	switch env {
	case "cloudDev":
		cfg = &mysql.Config{
			Addr:   secrets.CloudSQLDev.Address,
			User:   secrets.CloudSQLDev.Username,
			Passwd: secrets.CloudSQLDev.Password,
			DBName: secrets.CloudSQLDev.DBName,
		}
	case "local":
		cfg = &mysql.Config{
			Addr:   secrets.LocalSQLDev.Address,
			User:   secrets.LocalSQLDev.Username,
			Passwd: secrets.LocalSQLDev.Password,
			DBName: secrets.LocalSQLDev.DBName,
		}
	case "prod":
		err = fmt.Errorf("production not implmenented")
	default:
		err = fmt.Errorf("env %q not valid, please use a valid one", env)
	}
	return cfg, err
}

func getSQLConfig(fileName string) (*config, error) {
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
