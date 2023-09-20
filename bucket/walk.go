package bucket

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/utils/where"
	"github.com/charmbracelet/log"
)

// recursively walk a bucket dir and extract all of its manifest file
// names and their paths
func Walk(bucket string) []ManifestFile {
	manifests := []ManifestFile{}
	bucketPath := filepath.Join(where.Buckets(), bucket)
	if !utils.PathExists(bucketPath) {
		log.Info("Did not find bucket", "path", bucketPath)
		return manifests
	}
	filepath.WalkDir(bucketPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		b := ManifestFile{
			Name: strings.TrimSuffix(d.Name(), ".json"),
			Path: path,
		}
		manifests = append(manifests, b)
		return nil
	})
	return manifests
}

type ManifestFile struct {
	Name string
	Path string
}
