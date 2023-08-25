package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/Yakiyo/scoob/pkg/bucket"
	"github.com/Yakiyo/scoob/utils"
	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var bucketCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Manage local buckets",
}

var bucketAdd = &cobra.Command{
	Use:     "add",
	Short:   "Install a bucket locally",
	Aliases: []string{"install"},
	Example: `scoob bucket add main
scoob bucket add somebucket https://github.com/some/bucket`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {
		if !utils.GitExists() {
			utils.Error("Unable to find git in path. Git is required for installing buckets")
		}
		var buckName, buckUrl string
		if len(args) == 2 {
			buckName = args[0]
			buckUrl = args[1]
		} else {
			buckName = args[0]
			log.Info("Attempting to find bucket url from known buckets", "bucket", buckName)
			buckUrl = bucket.KnownBuckets[buckName]
			if buckUrl == "" {
				utils.Error("Not a known bucket. Please provide a url to the bucket.", "bucket", buckName)
				return
			}
			log.Info("Resolved bucket url from known buckets", "url", buckUrl)
		}
		if bucket.Exists(buckName) {
			utils.Error("Bucket already exists. Consider uninstalling it first and then reinstalling", "bucket", buckName)
		}
		bucketPath := bucket.GetPath(buckName)
		if parent := filepath.Dir(bucketPath); parent != "" && !utils.PathExists(parent) {
			err := os.MkdirAll(parent, os.ModePerm)
			if err != nil {
				utils.Error("Unable to create parent dir of buckets", "err", err)
			}
		}
		code, err := utils.RunGit(true, "clone", buckUrl, bucketPath)
		if err != nil {
			utils.Error("Failed to clone bucket, error running git", "exitCode", code, "err", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 4, ' ', 0)
		fmt.Println("Successfully installed bucket")
		fmt.Fprintln(w, color.BlueString("Name")+"\t"+buckName)
		fmt.Fprintln(w, color.BlueString("Source")+"\t"+buckUrl)
		fmt.Fprintln(w, color.BlueString("Path")+"\t"+bucketPath)
		w.Flush()
	},
}

var bucketRemove = &cobra.Command{
	Use:     "remove",
	Short:   "Uninstall a local bucket",
	Example: "scoob bucket remove main",
	Aliases: []string{"rm"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		buck := args[0]
		if !bucket.Exists(buck) {
			utils.Error("No bucket with that name is installed", "bucket", buck)
		}
		buckPath := bucket.GetPath(buck)
		if err := os.RemoveAll(buckPath); err != nil {
			utils.Error("Failed to uninstall bucket", "bucket", buck, "err", err)
		}
		fmt.Println("Successfully removed bucket", color.CyanString(buck))
	},
}

// TODO: bucket listing should be made better
// print bucket name with bucket source (git url)
var bucketList = &cobra.Command{
	Use:     "list",
	Short:   "List all locally installed buckets",
	Example: "scoob bucket list",
	Aliases: []string{"ls"},
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		l, err := bucket.List()
		if err != nil {
			utils.Error("Unable to read buckets directory", "err", err)
		}
		if len(l) < 1 {
			fmt.Println("No buckets installed locally")
			return
		}
		for _, i := range l {
			fmt.Println(i)
		}

	},
}

func init() {
	bucketCmd.AddCommand(bucketAdd)
	bucketCmd.AddCommand(bucketRemove)
	bucketCmd.AddCommand(bucketList)
	rootCmd.AddCommand(bucketCmd)
}
