package runtimes

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetWorkingDir get working dir
func GetWorkingDir() string {
	dir := GetExecutablePathByBuild()
	if strings.Contains(dir, GetTempDir()) {
		return GetExecutablePathByCaller()
	}
	return dir
}

// GetTempDir  get system temp dir
func GetTempDir() string {
	dir := os.Getenv("TEMP")
	if dir == "" {
		dir = os.Getenv("TMP")
	}
	res, _ := filepath.EvalSymlinks(dir)
	return res
}

// GetExecutablePathByBuild get current executable path by build
func GetExecutablePathByBuild() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// GetExecutablePathByCaller get current executable path by caller
func GetExecutablePathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(3)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}
