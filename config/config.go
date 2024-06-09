package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	ServerInfo struct {
		Port string
		Info string
	}

	MongoDB struct {
		Uri string
		Db  string
	}
}

func NewConfig(path string) *Config {
	c := new(Config)

	if f, err := os.Open(path); err != nil {
		panic(err)
	} else if err := toml.NewDecoder(f).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
