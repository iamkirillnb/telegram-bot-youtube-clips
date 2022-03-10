package internal

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type TelegramBot struct {
	Token string `yaml:"token"`
}

type Youtube struct {
	Url          string `yaml:"url"`
	ApiKey       string `yaml:"apikey"`
}

type Config struct {
	Telegram TelegramBot `yaml:"telegram"`
	Youtube  Youtube     `yaml:"youtube"`
}

var instance *Config
var once sync.Once

func GetConfig(path string) *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(path, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Println(help)
			log.Fatal(err)
		}
	})
	return instance
}
