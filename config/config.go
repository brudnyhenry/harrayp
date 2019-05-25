package config

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/viper"
)

var (
	URL      = ""
	Login    = ""
	Password = ""
)

// LoadConfigFile reads the specified config file
func LoadConfigFile() error {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".harrayp" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetConfigName(".harrayp")

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	// Set values. Config file will override commandline
	Login = viper.GetString("login")
	Password = viper.GetString("password")
	URL = viper.GetString("url")

	return nil
}
