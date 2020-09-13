package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// FindIcon find the icon location
func FindIcon() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(path.Join(exPath, "verto_desktop.png"))
	return path.Join(exPath, "verto_desktop.png")
}
