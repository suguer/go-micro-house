package utils

import (
	"path/filepath"
	"runtime"
)

func WorkDir() string {
	_, workDir, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filepath.Dir(workDir)))
}
