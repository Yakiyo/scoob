package bucket

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/where"
	"github.com/charmbracelet/log"
)

// recursively walk a bucket dir and extract all of its manifest file
// names and their paths
func Walk(bucket string) []BucketDir {
	manifests := []BucketDir{}
	bucketPath := filepath.Join(where.Buckets(), bucket)
	if !utils.PathExists(bucketPath) {
		log.Info("Did not find bucket", "path", bucketPath)
		return []BucketDir{}
	}
	filepath.WalkDir(bucketPath, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}
		b := BucketDir{
			Name: strings.TrimSuffix(d.Name(), ".json"),
			Path: path,
		}
		manifests = append(manifests, b)
		return nil
	})
	return manifests
}
