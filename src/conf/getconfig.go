package conf

import (
	"github.com/spf13/viper"
	. "github.com/aceberg/HomeLists/models"
)

const configPath = "/data/homelists/config"

func GetConfig() (config Conf) {
	viper.SetDefault("DB_PATH", "/data/homelists/sqlite.db")
	viper.SetDefault("GUI_IP", "0.0.0.0")
	viper.SetDefault("GUI_PORT", "8842")
	viper.SetDefault("THEME", "minty")

    viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
    viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DbPath = viper.Get("DB_PATH").(string)
	config.GuiIP = viper.Get("GUI_IP").(string)
	config.GuiPort = viper.Get("GUI_PORT").(string)
	config.Theme = viper.Get("THEME").(string)

	return config
}

// func write_config() {
// 	viper.SetConfigFile(configPath)
// 	viper.SetConfigType("env")
// 	viper.Set("THEME", AppConfig.Theme)
// 	viper.WriteConfig()
// }