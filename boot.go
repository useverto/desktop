package main

import (
	"os"
	"path"
)

// DetectBootstrap detect if the bootstrap process has already been done
func DetectBootstrap() bool {
	tmpVendorPath := path.Join(os.TempDir(), "./vendor/astilectron")
	if _, err := os.Stat(tmpVendorPath); err == nil {
		return true
	}
	return false
}
