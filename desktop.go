package main

import (
	"log"

	"github.com/mitchellh/go-homedir"
	_ "github.com/useverto/desktop/bundle"
	"github.com/useverto/desktop/webview"
)

func main() {
	log.Println("Starting thread loop")
	// init
	debug := true
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

	// create webview instance
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Verto")
	w.SetSize(2000, 2000, webview.HintNone)
	// bind methods
	w.Bind("quit", func() {
		w.Terminate()
	})

	// Render view
	w.Navigate("http://localhost:8000/")

	// Run webview
	w.Run()
}
