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
			Addr:   c.CloudSQLDev.Address,
			User:   c.CloudSQLDev.Username,
			Passwd: c.CloudSQLDev.Password,
			DBName: c.CloudSQLDev.DBName,
		}
	case "local":
		cfg = &mysql.Config{
			Addr:   c.LocalSQLDev.Address,
			User:   c.LocalSQLDev.Username,
			Passwd: c.LocalSQLDev.Password,
			DBName: c.LocalSQLDev.DBName,
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
