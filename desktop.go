package main

import (
	"log"

	"github.com/useverto/desktop/webview"

	_ "github.com/useverto/desktop/bundle"
)

func main() {
	log.Println("Starting thread loop")
	// init
	debug := true

	if NeedsUpgrade() {
		log.Printf("%s is not latest, upgrading...", version)
		DownloadRelease(".")
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
