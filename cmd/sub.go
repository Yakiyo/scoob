package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage local buckets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sub called")
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
