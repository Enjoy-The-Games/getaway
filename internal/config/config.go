package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type Config struct {
	IsDebug bool `yaml:"is_debug" env-required="true"`
	Listen  struct {
		Port string `yaml:"port" env-default="8080"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

func GetConfig(configName string) *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig(configName, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatal(help)
		}
	})
	return instance
}
