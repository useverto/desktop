package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/mitchellh/go-homedir"
	_ "github.com/useverto/desktop/bundle"
)

func iconName() string {
	if runtime.GOOS == "darwin" {
		return "verto_desktop_mac"
	}
	return "verto_desktop"
}

func main() {
	log.Println("Starting thread loop")

	if !FindIcon() {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		DownloadFile("https://github.com/useverto/desktop/raw/master/assets/"+iconName()+".png", path.Join(exPath, "verto_desktop.png"))
	}

	_, err := setupWatcher()

	if err != nil {
		// do something sensible
		log.Fatal(err)
	}

	// upgrade the desktop version
	if NeedsUpgrade() {
		log.Printf("%s is not latest, upgrading...", version)
		// determine download location
		downloadLoc, _ := homedir.Expand(`~/.verto_desktop`)
		// download latest release from github
		DownloadRelease(downloadLoc)
		// unzip the release zip
		NewUnzip(downloadLoc, downloadLoc).Extract()

	}

	// start the server with website source
	Loadview()
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())
	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Verto Desktop",
		BaseDirectoryPath: "verto_desktop",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: boot failed: %w", err))
	}
	defer a.Close()

	// Handle signals
	a.HandleSignals()

	// Start
	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: boot failed: %w", err))
	}

	// New window
	var w *astilectron.Window
	if w, err = a.NewWindow("http://localhost:8000", &astilectron.WindowOptions{
		Title:  astikit.StrPtr("Verto"),
		Height: astikit.IntPtr(3000),
		Width:  astikit.IntPtr(3000),
		Icon:   astikit.StrPtr(LoadIcon()),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	// Create windows
	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// Blocking pattern
	a.Wait()
}
