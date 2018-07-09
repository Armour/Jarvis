// Package config contains functions that load, parse and return jarvis configs.
package config

import (
	"github.com/armour/jarvis/internal/pkg/utils"
	"github.com/spf13/viper"
)

var (
	jarvisConfig = viper.New()
)

// LoadConfig loads jarvis configs.
func LoadConfig() {
	jarvisConfig.SetDefault("author", "Author")
	jarvisConfig.SetDefault("email", "your.email@email.com")
	jarvisConfig.SetDefault("githubUser", "GithubUser")

	jarvisConfig.SetConfigName("jarvis") // name of config file (without extension)
	jarvisConfig.AddConfigPath("$HOME")  // path to look for the config file in
	jarvisConfig.AddConfigPath(".")      // optionally look for config in the working directory
	if err := jarvisConfig.ReadInConfig(); err != nil {
		utils.ExitOnError(err)
	}
}

// GetConfigByField gets jarvis config by field name.
func GetConfigByField(field string) string {
	return jarvisConfig.GetString(field)
}
