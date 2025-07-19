package config

import (
	"os"

	"github.com/spf13/viper"
)

var (
	GitHubAPIToken   string
	DeepseekAPIToken string
)

func init() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	GitHubAPIToken = viper.GetString("GithubAPIToken")
	DeepseekAPIToken = viper.GetString("DeepseekAPIToken")
}
