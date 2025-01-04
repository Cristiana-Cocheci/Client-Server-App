package src

import (
	"client-server/config"
	"os"
	"path/filepath"
)

var REQUEST_MAP = map[string]func([]string) string{
	"Blank": func([]string) string { return "" },
	"ex1": func(args []string) string {
		return MixedLetters(args)
	},
	"ex3": func(args []string) string {
		return SolveEx3(args)
	},
	"ex5": func(args []string) string {
		return SolveEx5(args)
	},
	"ex6": func(args []string) string {
		return SolveEx6(args)
	},
	"ex7": func(args []string) string {
		return SolveEx7(args[0])
	},
	"ex8": func(args []string) string {
		return SolveEx8(args)
	},
	"ex12": func(args []string) string {
		return SolveEx12(args)
	},
	"map_reduce_6": func(args []string) string {
		return mapReduce6(args)
	},
	"exit": func([]string) string { return "exit" },
}

var WD, _ = os.Getwd()

var ConfigPath = filepath.Join(WD, "env.txt")
var conf = config.LoadConfig(ConfigPath)

var ExampleClientsNumber = 5
