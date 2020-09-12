package main

import (
	"fmt"
	"log"

	"github.com/mitchellh/go-homedir"
	_ "github.com/useverto/desktop/bundle"
	"github.com/zserge/lorca"
)

func main() {
	log.Println("Starting thread loop")
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

	ui, err := lorca.New("http://localhost:8000/", "", 3000, 3000)
	if err != nil {
		fmt.Println(err)
	}
	defer ui.Close()

	// Wait for the browser window to be closed
	<-ui.Done()
}
