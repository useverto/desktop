run:
	go run .

build:
	go build .

bundle-mac:
	# creates a MACOS .app bundle for the application
	# NOTE: bundle the app with using node-appdmg for producing DMG images
	sh ./tools/create_mac_app.sh

bundle-linux:
	# creates a debian using go-deb
	sh ./tools/create_deb.sh

bundle-win:
	# msitools is restricted for use only on Linux and MacOS -> so is Make
	# on Windows, use Wix Toolset 3 for building an MSI
	# See, .github/workflows/ci.yml for building with Wix

	# Install msitools using `sudo apt-get install wixl`
	wixl -v verto.wxs