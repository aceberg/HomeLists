package conf

import (
	"github.com/spf13/viper"

	"github.com/aceberg/HomeLists/internal/models"
)

const configPath = "/data/homelists/config"

func GetConfig() (config models.Conf) {
	viper.SetDefault("DB_PATH", "/data/homelists/sqlite.db")
	viper.SetDefault("GUI_IP", "0.0.0.0")
	viper.SetDefault("GUI_PORT", "8842")
	viper.SetDefault("THEME", "superhero")
	viper.SetDefault("COLOR", "light")
	viper.SetDefault("NODEPATH", "")

	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.ReadInConfig()

	viper.AutomaticEnv() // Get ENVIRONMENT variables

	config.DbPath = viper.Get("DB_PATH").(string)
	config.GuiIP = viper.Get("GUI_IP").(string)
	config.GuiPort = viper.Get("GUI_PORT").(string)
	config.Theme = viper.Get("THEME").(string)
	config.Color = viper.Get("COLOR").(string)
	config.NodePath = viper.Get("NODEPATH").(string)

	return config
}

func WriteConfig(theme string) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("env")
	viper.Set("THEME", theme)
	viper.WriteConfig()
}
