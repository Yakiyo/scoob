package cmd

import (
	"fmt"
	"os"

	"github.com/Yakiyo/scoob/pkg/bucket"
	"github.com/charmbracelet/log"
	"github.com/samber/lo"
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

var bucketList = &cobra.Command{
	Use:     "list",
	Short:   "List all locally installed buckets",
	Example: "scoob bucket list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		l, err := bucket.ListBucket()
		if err != nil {
			log.Error("Unable to read buckets directory", "err", err)
			os.Exit(1)
		}
		if len(l) < 1 {
			fmt.Println("No buckets installed locally")
			return
		}
		lo.ForEach(l, func(it string, _ int) {
			fmt.Println(l)
		})

	},
}

func init() {
	bucketCmd.AddCommand(bucketAdd)
	bucketCmd.AddCommand(bucketList)
	rootCmd.AddCommand(bucketCmd)
}
