package main

import (
		"github.com/enocPierre/api-users/configs"
)

func main() {
	config, _:= configs.LoadConfig(".")
	println(config.DBDriver)
}