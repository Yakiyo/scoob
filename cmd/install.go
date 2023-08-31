package cmd

import "github.com/spf13/cobra"

var installCmd = &cobra.Command{
	Use: "install",
	Short: "Install an app",
	Args: cobra.ExactArgs(1),
	Aliases: []string{"add"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}

func init() {
	f := installCmd.Flags()
	f.BoolP("global", "g", false, "Install globally (unimplemented)")
	f.BoolP("independent", "i", false, "Don't install dependencies automatically")
	f.BoolP("no-cache", "k", false, "Don't use download cache")
	f.BoolP("no-update-scoop", "u", false, "Don't update Scoop before installing if it's outdated")
	f.BoolP("skip", "s", false, "Skip hash validation (use with caution!)")
	f.StringP("arch", "a", "", "Use the specified architecture, if the app supports it (32bit|64bit|arm64)")
	rootCmd.AddCommand(installCmd)
}