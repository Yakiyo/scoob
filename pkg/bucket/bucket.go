package bucket

import (
	"io/fs"
	"os"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/where"
	"github.com/samber/lo"
)

// list all local buckets
func ListBucket() ([]string, error) {
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
