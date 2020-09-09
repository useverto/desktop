<p align="center">
  <a href="https://verto.exchange">
    <img src="https://raw.githubusercontent.com/useverto/design/master/logo/logo_light.svg" alt="Verto logo (light version)" width="110" />
  </a>

  <h3 align="center">Verto Desktop</h3>

  <p align="center">
    All of verto's website in a tiny desktop app
  </p>

  <p align="center">
    <img src="https://github.com/useverto/desktop/workflows/ci/badge.svg" alt="Fancy CI badge" />
  </p>

</p>

## About

This repository contains all of the necessary code for Verto's desktop app.

You can access the code for our website [here](https://github.com/useverto/verto).

> Important Notice: Verto is in its Alpha stage. If you have a suggestion, idea, or find a bug, please report it! The Verto team will not be held accountable for any funds lost.

## Building from source

### Prerequisite

Before building the desktop application, you will need `go` installed on your machine.

You can install `statik` via Go CLI

 ```shell script
 go get github.com/rakyll/statik
 ```

**Linux**

Run the below command for your specific platform to install the Linux requirements.

```sh
# Ubuntu
sudo apt-get install libwebkit2gtk-4.0-dev
# Fedora
sudo dnf install webkit2gtk3-devel.x86_64
```
### Building

```sh
# build the website
make web
# build the desktop application
make build
```
## License

The code contained within this repository is licensed under the MIT license.
See [`./LICENSE`](./LICENSE) for more information.
