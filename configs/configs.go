package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Configs struct {
}

func NewConfigs() *Configs {
	c := new(Configs)
	return c
}

func (c *Configs) Init() (error) {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	env := os.Getenv("APP_ENV")
	if "" == env {
		env = "development"
	}

	return nil
}

func (c *Configs) Get(key string) (string) {
	env := os.Getenv(key)
	return env
}