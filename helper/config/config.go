package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name string `toml:"name"`
}

func FetchConfig(path string) Config {
	var config Config

	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		// check if an the config file has an error in it:
		log.Println(`Error in config file "` + path + `"`)
	}

	return config
}

func main() {
	fmt.Println(FetchConfig("helper/config.toml").Name)
}
