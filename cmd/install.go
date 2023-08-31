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
	rootCmd.AddCommand(installCmd)
}