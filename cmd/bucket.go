package cmd

import (
	"fmt"

	"github.com/Yakiyo/scoob/pkg/bucket"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage local buckets",
}

var bucketAdd = &cobra.Command{
	Use:   "add",
	Short: "Install a bucket locally",
	Example: `scoob bucket add main
scoob bucket add somebucket https://github.com/some/bucket`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		var buckName, buckUrl string
		if len(args) == 2 {
			buckName = args[0]
			buckUrl = args[1]
		} else {
			buckName = args[0]
			log.Info("Attempting to find bucket url from known buckets", "bucket", buckName)
			buckUrl = bucket.KnownBuckets[buckName]
			if buckUrl == "" {
				log.Error("Not a known bucket. Please provide a url to the bucket.", "bucket", buckName)
				return
			}
		}
		fmt.Println(buckName, buckUrl)
	},
}

func init() {
	bucketCmd.AddCommand(bucketAdd)
	rootCmd.AddCommand(bucketCmd)
}
