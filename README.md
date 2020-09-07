## Verto Desktop App

Desktop App for Verto

### Building from source

## Build static files
```sh
go get github.com/rakyll/statik
yarn build:web
statik -src=./views
```

#### Linux

```sh
# Ubuntu
sudo apt-get install libwebkit2gtk-4.0-dev
# Fedora
sudo dnf install webkit2gtk3-devel.x86_64
```