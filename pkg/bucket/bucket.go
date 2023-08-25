package bucket

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/where"
	"github.com/samber/lo"
)

// list all local buckets
func List() ([]string, error) {
	bucketPath := where.Buckets()
	if !utils.PathExists(bucketPath) {
		return []string{}, nil
	}
	entries, err := os.ReadDir(bucketPath)
	if err != nil {
		return []string{}, err
	}
	return lo.FilterMap(entries, func(f fs.DirEntry, _ int) (string, bool) {
		if f.IsDir() {
			return f.Name(), true
		}
		return "", false
	}), nil
}

// get the path to a bucket
func GetPath(name string) string {
	return filepath.Join(where.Buckets(), name)
}

// wether a bucket with that name exists
func Exists(name string) bool {
	return utils.PathExists(GetPath(name))
}
