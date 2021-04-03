package path

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

func GetRootDirectory() string {
	_, b, _, _ := runtime.Caller(0)
	path := filepath.Dir(b)
	sliceOfPath := strings.Split(path, "/") // old /  window ใช้ \
	fmt.Println(sliceOfPath)
	return strings.Join(sliceOfPath[:len(sliceOfPath)-3], "/")
}
