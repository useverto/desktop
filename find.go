package main

import (
	"os"
	"path"
	"path/filepath"
)

// LoadIcon find the icon location
func LoadIcon() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return path.Join(exPath, "verto_desktop.png")
}

// FindIcon check if icon exists
func FindIcon() bool {
	if _, err := os.Stat("file-exists.go"); err == nil {
		return true
	}
	return false
}
