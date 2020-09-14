package main

import (
	"os"
)

// FindIcon check if icon exists
func FindIcon(iconPath string) bool {
	if _, err := os.Stat(iconPath); err == nil {
		return true
	}
	return false
}
