package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/scoob/config"
	logger "github.com/Yakiyo/scoob/log"
	"github.com/Yakiyo/scoob/meta"
	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/where"
	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   meta.AppName,
	Short: "Drop-in replacement cli for scoop",
	Long: `Scoob is a rewrite of scoop in Go.
	
It aims to be a faster alternative to scoop, with the same features.
For any queries, issues or bug reports, please visit https://github.com/Yakiyo/scoob`,
	Version: meta.Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		config.BindFlags(cmd)

		if cfg := utils.Must(cmd.Flags().GetString("config")); cfg != "" {
			// when passed the `--config/-c` flag, set viper
			// to explicitly read from that file, this will
			// error if the config file does not exist
			viper.SetConfigFile(cfg)
		}

		// read config
		config.Read()

		logger.SetLevel(viper.GetString("log_level"))
		utils.SetColor(viper.GetString("color"))
		where.SetRoot(viper.GetString("root_dir"))
		log.Debug(viper.AllSettings())
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate(func() string {
		return fmt.Sprintf("Scoob version %v - ", color.BlueString(meta.Version)) +
			fmt.Sprintf("Released at %v - ", color.BlueString(meta.BuiltAt)) +
			fmt.Sprintf("(Revision %v)\n", color.BlueString(meta.Revision))
	}())

	cc.Init(&cc.Config{
		RootCmd:         rootCmd,
		Headings:        cc.HiCyan + cc.Bold + cc.Underline,
		Commands:        cc.HiYellow + cc.Bold,
		Example:         cc.Bold,
		ExecName:        cc.Bold,
		Flags:           cc.Bold,
		FlagsDataType:   cc.Italic + cc.HiBlue,
		NoExtraNewlines: true,
	})

	f := rootCmd.PersistentFlags()
	// allow users to set custom config files
	f.StringP("config", "c", "", "Path to config file")
	// dont mention debug level, usually users dont need, only on the dev side
	f.String("log-level", "", "Set log level [info, warn, error, fatal]")
	f.String("color", "", "Set color output [always, auto, never]")
}
