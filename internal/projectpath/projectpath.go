package projectpath

import (
	"path/filepath"
	"runtime"
)

func ProjectPath() string {
	_, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	root := filepath.Join(filepath.Dir(b), "../..")
	return root
}
