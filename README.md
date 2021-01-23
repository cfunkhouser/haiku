# Haiku

Simple generator of haiku-style names for whatever your naming needs.

[![CircleCI](https://img.shields.io/circleci/build/github/cfunkhouser/haiku/main)](https://app.circleci.com/pipelines/github/cfunkhouser/haiku)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cfunkhouser/haiku)
[![Go Reference](https://pkg.go.dev/badge/github.com/cfunkhouser/haiku.svg)](https://pkg.go.dev/github.com/cfunkhouser/haiku)


There is a library, and there is a CLI tool. They are both so simple, this
README is insulting.


## CLI Example

Fine, I'll show you.

```console
$ go install ./cmd/haiku
$ haiku
tubby-endothelium
$ haiku -number
chippy-oarfish-150
```

## Word Lists

The words used in this tool are sourced from [Princeton
WordNet](https://wordnet.princeton.edu/).