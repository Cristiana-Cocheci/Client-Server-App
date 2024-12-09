package config

import (
	"fmt"
	"os"
)

func TryError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func PrintError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
