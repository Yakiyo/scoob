package cmd

import "github.com/spf13/cobra"

var installCmd = &cobra.Command{
	Use: "install",
	Short: "Install an app",
	Long: `Install apps via scoob
Help: e.g. The usual way to install an app (uses your local 'buckets'):
    scoob install git

To install a different version of the app
(note that this will auto-generate the manifest using current version):
    scoob install gh@2.7.0

To install an app from a manifest at a URL:
    scoob install https://raw.githubusercontent.com/ScoopInstaller/Main/master/bucket/runat.json

To install an app from a manifest on your computer
    scoob install \path\to\app.json`,
	Args: cobra.ExactArgs(1), // TODO: support installing multiple apps in the future
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