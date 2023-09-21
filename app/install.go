package app

import (
	"fmt"
	"time"

	"github.com/Yakiyo/scoob/manifest"
	"github.com/charmbracelet/log"
	jsoniter "github.com/json-iterator/go"
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
	_, err := get_arch(m, architecture)
	if err != nil {
		return err
	}
	if viper.GetBool("show_manifest") {
		j := lo.Must(jsoniter.MarshalIndent(m, "", " "))
		fmt.Printf("%v\n", string(j))
	}
	fmt.Printf("Installing %v (%v) [%v]", app.Name, app.Version, architecture)
	if app.Bucket != "" {
		fmt.Printf(" from %v bucket", app.Bucket)
	}
	fmt.Println()

	return nil
}

func get_arch(m manifest.Manifest, arch string) (manifest.Arch, error) {
	archs := m.Architecture
	r, ok := archs[arch]
	if !ok {
		r = manifest.Arch{
			Bin:         m.Bin,
			Installer:   m.Installer,
			Uninstaller: m.Uninstaller,
			Hash:        m.Hash,
			Url:         m.Url,
			Extract_dir: m.Extract_dir,
			Extract_to:  m.Extract_to,
			Shortcuts:   m.Shortcuts,
		}
	}
	return r, nil
}
