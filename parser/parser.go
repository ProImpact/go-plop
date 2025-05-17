// Package parser provides all the functions for parsing components
package parser

import (
	"fmt"

	"github.com/spf13/viper"
)

func Test() {
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			fmt.Println("config founded")
		} else {
			// Config file was found but another error was produced
		}
	}

	// Config file found and successfully parsed
}
