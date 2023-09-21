package config

import (
	"path/filepath"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/utils/meta"
	"github.com/Yakiyo/scoob/utils/where"
	"github.com/charmbracelet/log"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/samber/lo"
	v "github.com/spf13/viper"
)

func init() {
	// config file named after the app
	v.SetConfigName(meta.AppName)
	v.SetConfigType("toml")

	// search in default directory - ~/.config/
	v.AddConfigPath(filepath.Join(lo.Must(homedir.Dir()), ".config"))

	// add default values here
	v.SetDefault("log_level", "warn")
	v.SetDefault("color", "auto")
	v.SetDefault("root_path", where.RootDir())
	v.SetDefault("default_architecture", lo.Must(utils.GetDefaultArch()))
}

// read config file
func Read() {

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(v.ConfigFileNotFoundError); ok {
			log.Info("Missing log file in default location. Using defaults")
		} else {
			log.Fatal("Error reading config", "err", err)
		}
	}
}
