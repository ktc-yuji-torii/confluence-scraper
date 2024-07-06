package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	BaseURL      string
	Username     string
	APIToken     string
	ParentPageID string
	Debug        bool
}

func LoadConfig(cmd *cobra.Command) (*Config, error) {
	// コマンドライン引数をViperにバインド
	viper.BindPFlag("baseURL", cmd.PersistentFlags().Lookup("baseURL"))
	viper.BindPFlag("username", cmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("apiToken", cmd.PersistentFlags().Lookup("apiToken"))
	viper.BindPFlag("parentPageID", cmd.PersistentFlags().Lookup("parentPageID"))
	viper.BindPFlag("debug", cmd.PersistentFlags().Lookup("debug"))

	// Viperを使用して設定値を取得
	config := &Config{
		BaseURL:      viper.GetString("baseURL"),
		Username:     viper.GetString("username"),
		APIToken:     viper.GetString("apiToken"),
		ParentPageID: viper.GetString("parentPageID"),
		Debug:        viper.GetBool("debug"),
	}

	return config, nil
}
