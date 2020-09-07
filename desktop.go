package main

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
	"github.com/webview/webview"

	_ "github.com/useverto/desktop/statik"
)

func main() {
	log.Println("Starting thread loop")
	// init
	debug := true
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

// Loadview load view
func Loadview() {

	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	// Serve the contents over HTTP.
	http.Handle("/", http.StripPrefix("/", http.FileServer(statikFS)))
	go func() {
		http.ListenAndServe(":8000", nil)
	}()
}
