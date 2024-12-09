package config

import (
	"fmt"
	"os"
	"strconv"
	str "strings"
)

type Config struct {
	ArrayLength   int64
	ClientsNumber int64
	ReadFromFile  bool
}

func read_file(path string) string {
	fstr, err := os.ReadFile(path)
	TryError(err)
	return string(fstr)
}

func addValueToConfig(config *Config, arg string, val string) {
	err := error(nil)
	switch arg {
	case "ArrayLength":
		config.ArrayLength, err = strconv.ParseInt(val, 10, 64)
	case "ClientsNumber":
		config.ClientsNumber, err = strconv.ParseInt(val, 10, 64)
	case "ReadFromFile":
		config.ReadFromFile, err = strconv.ParseBool(val)
	}
	TryError(err)
}

func LoadConfig(path string) (config Config) {
	fstr := read_file(path)
	lineArray := str.Split(string(fstr), "\n")
	config = Config{}

	for _, line := range lineArray {
		arg, val := str.Split(line, "=")[0], str.Split(line, "=")[1]
		addValueToConfig(&config, arg, val)
	}
	fmt.Printf("Config Loaded\n")
	return config
}

func LoadRequests(path string) []string {
	fstr := read_file(path)
	return str.Split(string(fstr), "\n")
}