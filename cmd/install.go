package cmd

import (
	"errors"
	"fmt"

	"github.com/Yakiyo/scoob/app"
	"github.com/Yakiyo/scoob/bucket"
	"github.com/Yakiyo/scoob/utils"
	"github.com/charmbracelet/log"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var installCmd = &cobra.Command{
	Use:   "install",
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
	Args:         cobra.ExactArgs(1), // TODO: support installing multiple apps in the future
	Aliases:      []string{"add"},
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var arch string
		if fv := lo.Must1(cmd.Flags().GetString("arch")); fv != "" {
			arch = fv
		} else {
			arch = viper.GetString("default_architecture")
		}
		_ = arch

		//lint:ignore SA9003 todo
		if !lo.Must1(cmd.Flags().GetBool("no-update-scoop")) {
			// TODO: update scoop
		}
		app, err := app.ParseApp(args[0])
		if err != nil {
			return err
		}
		if app.Name == "" {
			return errors.New("Invalid app name. Did not match app syntax [bucket/]app[@version]")
		}
		buckets := []string{}
		// if a specific bucket was given, use it
		if app.Bucket != "" {
			p := bucket.GetPath(app.Bucket)
			if !utils.PathExists(p) {
				return fmt.Errorf("No bucket with name %v installed locally", app.Bucket)
			}
			buckets = append(buckets, p)
		} else {
			bdirs, err := bucket.List()
			if err != nil {
				return err
			}
			buckets = append(buckets, lo.Map[bucket.BucketDir, string](bdirs, func(item bucket.BucketDir, _ int) string {
				return item.Path
			})...)
		}
		var manifests []bucket.ManifestFile
		for _, buck := range buckets {
			manifests = bucket.Walk(buck)
		}
		manifest, found := lo.Find(manifests, func(i bucket.ManifestFile) bool { return i.Name == app.Name })
		if !found {
			return fmt.Errorf("Unable to find manifest with name %v in local buckets", app.Name)
		}
		log.Info("Found manifest", "name", manifest.Name, "file", manifest.Path)
		fmt.Println(manifest)
		if app.Version != "" {
			log.Warn("Handling non-latest version is not implemented yet. Installing latest version")
		}

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
