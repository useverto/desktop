package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Dev43/arweave-go/wallet"
	"github.com/mitchellh/go-homedir"
	"github.com/ncruces/zenity"
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
	w.Init(`
	function addStyle(styleString) {
		const style = document.createElement('style');
		style.textContent = styleString;
		document.head.append(style);
	}
	addStyle("* { -webkit-font-smoothing: antialiased; -webkit-text-stroke: 0.5px; }");
	let x = setInterval(() => assignFileDialog(), 200);
	async function assignFileDialog() {
		if(window.location.pathname.startsWith("/login")) {
			clearInterval(x)
			let val = await window.native_file_dialog();
			let addr = await window.native_wallet_addr(val);
			localStorage.setItem("keyfile", val);
			localStorage.setItem("address", addr);
			window.location.href = "/app"
		}
	}
	`)
	// open a native file dialog (only mac) and get file content
	w.Bind("native_file_dialog", func() string {
		file, err := zenity.SelectFile()
		if err != nil {
			fmt.Println("File reading error", err)
			return ""
		}
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println("File reading error", err)
			return ""
		}
		return string(data)
	})
	w.Bind("native_wallet_addr", func(keyfile string) string {
		// create a new wallet instance
		w := wallet.NewWallet()
		// extract the key from the wallet instance
		err = w.LoadKey([]byte(keyfile))
		if err != nil {
			fmt.Println("File reading error", err)
		}
		return w.Address()
	})
	// Render view
	w.Navigate("http://localhost:8000/")

	// Run webview
	w.Run()
}
