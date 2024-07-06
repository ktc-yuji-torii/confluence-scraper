package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL    string
	Username   string
	APIToken   string
	HomepageID string
	Debug      bool
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	// コマンドライン引数をViperにバインド
	viper.BindPFlag("baseURL", cmd.PersistentFlags().Lookup("baseURL"))
	viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("apiToken", cmd.PersistentFlags().Lookup("apiToken"))
	viper.BindPFlag("homepageID", cmd.PersistentFlags().Lookup("homepageID"))
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// Viperを使用して設定値を取得
	config := &Config{
		BaseURL:    viper.GetString("baseURL"),
		Username:   viper.GetString("username"),
		APIToken:   viper.GetString("apiToken"),
		HomepageID: viper.GetString("homepageID"),
		Debug:      viper.GetBool("debug"),
	}

	return config, nil
}
