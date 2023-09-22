package app

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Yakiyo/scoob/manifest"
	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/utils/where"
	"github.com/charmbracelet/log"
	json "github.com/json-iterator/go"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

// install an app
func Install(app ParsedApp, m manifest.Manifest, architecture string, global bool, use_cache bool, check_hash bool) error {
	if !check_hash {
		log.Info("Skipping hash check")
	}
	if m.Version == "nightly" {
		log.Info("Installing nightly version. Hash check will be skipped")
		m.Version = fmt.Sprintf("nightly-%v", time.Now().Format("2006.01.02"))
		check_hash = false
	}
	arch, err := get_arch(m, architecture)
	if err != nil {
		return err
	}
	if viper.GetBool("show_manifest") {
		j := lo.Must(json.MarshalIndent(m, "", " "))
		fmt.Printf("Manifest: %v.json\n", app.Name)
		fmt.Printf("%v\n", string(j))
	}
	fmt.Printf("Installing %v (%v) [%v]", app.Name, app.Version, architecture)
	if app.Bucket != "" {
		fmt.Printf(" from %v bucket", app.Bucket)
	}
	fmt.Println()

	vdir := filepath.Join(where.Apps(), app.Name, m.Version)
	if err := utils.EnsureDir(vdir); err != nil {
		return utils.Anyhow("Failed to ensure version dir", err)
	}
	_, _ = url_fname(arch.Url[0])

	return nil
}
