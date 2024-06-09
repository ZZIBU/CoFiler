package main

import (
	"Core/config"
	"Core/init/app"
	"flag"
)

var envFlag = flag.String("config", "./env.toml", "env file not found")

func main() {
	flag.Parse()
	c := config.NewConfig(*envFlag)
	app.NewApp(c)
}
