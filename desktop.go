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
	"github.com/gen2brain/beeep"
)

func iconName() string {
	if runtime.GOOS == "darwin" {
		return ".icns"
	}
	return ".png"
}

func main() {
	log.Println("Starting thread loop")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	iconExt := iconName()
	iconPath := path.Join(exPath, "verto_desktop"+iconExt)
	if !FindIcon(iconPath) {
		DownloadFile("https://github.com/useverto/desktop/raw/master/assets/verto_desktop"+iconExt, iconPath)
	}

	_, err = setupWatcher()

	if err != nil {
		// do something sensible
		log.Fatal(err)
	}

	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())
	if !DetectBootstrap() {
		err := beeep.Alert("Verto Desktop", "Hang on for a few seconds while Verto initialises...", iconPath)
		if err != nil {
			fmt.Println("main: user alert failed: %w", err)
		}
	}
	// Create astilectron
	a, err := astilectron.New(l, astilectron.Options{
		AppName:            "Verto Desktop",
		BaseDirectoryPath:  os.TempDir(),
		AppIconDefaultPath: iconPath,
		AppIconDarwinPath:  iconPath,
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
	if w, err = a.NewWindow("https://verto.exchange/login", &astilectron.WindowOptions{
		Title:  astikit.StrPtr("Verto"),
		Height: astikit.IntPtr(3000),
		Width:  astikit.IntPtr(3000),
		Icon:   astikit.StrPtr(iconPath),
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
