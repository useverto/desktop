package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
)

func detectHostOS() string {
	if runtime.GOOS == "darwin" {
		return "macos"
	}
	return runtime.GOOS
}

// DownloadRelease download latest github release for useverto/desktop
func DownloadRelease(filepath string) (err error) {

	hostOS := detectHostOS()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get("https://github.com/useverto/trading-post/releases/latest/download/verto-x64-" + hostOS)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
