package bucket

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/Yakiyo/scoob/utils"
	"github.com/Yakiyo/scoob/utils/where"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

// list all local buckets
func List() ([]BucketDir, error) {
	bucketPath := where.Buckets()
	if !utils.PathExists(bucketPath) {
		return []BucketDir{}, nil
	}
	entries, err := os.ReadDir(bucketPath)
	if err != nil {
		return []BucketDir{}, err
	}
	// run it parallely for speeding it up
	return lop.Map(
		lo.Filter(entries, func(i fs.DirEntry, _ int) bool { return i.IsDir() }),
		func(f fs.DirEntry, _ int) BucketDir {
			return fromDir(f, filepath.Join(bucketPath, f.Name()))
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
