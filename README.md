
# Gobox

A simple CLI tool to manage and reuse Go packages efficiently.

Easily initialize Go projects with your favorite packages, reuse dependencies, and speed up your development workflow.

[![Release](https://github.com/MegeKaplan/gobox/actions/workflows/release.yaml/badge.svg)](https://github.com/MegeKaplan/gobox/actions)
[![Go Modules](https://img.shields.io/github/go-mod/go-version/MegeKaplan/gobox?logo=go)](https://github.com/MegeKaplan/gobox)
[![Go Report Card](https://goreportcard.com/badge/github.com/MegeKaplan/gobox)](https://goreportcard.com/report/github.com/MegeKaplan/gobox)
[![License: MIT](https://img.shields.io/github/license/MegeKaplan/gobox)](https://choosealicense.com/licenses/mit/)
[![Downloads](https://img.shields.io/github/downloads/MegeKaplan/gobox/total.svg)](https://github.com/MegeKaplan/gobox/releases)
[![Platform](https://img.shields.io/badge/Tested%20on-Linux-blue?logo=linux)](#)
[![GitHub last commit](https://img.shields.io/github/last-commit/MegeKaplan/gobox.svg)](https://github.com/MegeKaplan/gobox/commits/main)
[![Activity](https://img.shields.io/github/commit-activity/m/MegeKaplan/gobox)](https://github.com/MegeKaplan/gobox/graphs/commit-activity)
[![Repo Size](https://img.shields.io/github/repo-size/MegeKaplan/gobox)](https://github.com/MegeKaplan/gobox)
[![Lines of Code](https://tokei.rs/b1/github/MegeKaplan/gobox?category=lines)](https://github.com/MegeKaplan/gobox)

## Features

**Initialize a new Go project** by selecting from your saved package collection
```bash
gobox init [project_name]
```

**Download and save a Go package** to your local collection for future use
```bash
gobox get [package_name]
```

**Remove a saved package** from your local collection
```bash
gobox remove
```

**Display all packages** in your local collection
```bash
gobox list
```

## Installation

You can install gobox by downloading precompiled binary or by building it directly from source.

### Install on Linux (via Releases)

```bash
curl -L https://github.com/MegeKaplan/gobox/releases/download/v1.1.0/gobox -o gobox
chmod +x gobox
sudo mv gobox /usr/local/bin
```

### Build from Source

```bash
git clone https://github.com/MegeKaplan/gobox
cd gobox
go build -o ./bin/gobox
sudo cp ./bin/gobox /usr/local/bin
```

Now you can run gobox from anywhere in your terminal with `gobox`

## Contributing

Open source is built on community contributions, and you're welcome to join in!

Whether it’s fixing a bug, improving documentation, or adding a feature, your help is appreciated.

Just fork the repo, make your changes, and open a pull request.

Let’s build together!

## License

[MIT](https://choosealicense.com/licenses/mit/)

