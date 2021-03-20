package localenv

import (
	"path"
	"path/filepath"
	"runtime"
)

// RootDir returns the root directory of the application
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
