# tago [![Go Report Card](https://goreportcard.com/badge/github.com/FalcoSuessgott/tago)](https://goreportcard.com/badge/github.com/FalcoSuessgott/tago) 

<p align="center">
  <img src="demo.gif" />
</p>

`tago` lets you bump git tags using [semantic versioning](https://semver.org/).

# Features
* detecting and handling semver tags with or without "v"-prefix
* check if there are uncommited changes 
* creating initial tag if no tags exists
* add lightweight or annotated [tags](https://git-scm.com/book/en/v2/Git-Basics-Tagging)
* interactive user prompting
* automatable using cli params for scripting purposes
* automatically push after tag creation

# when to bump a version
`MAJOR.MINOR.PATCH`
* Major: when you make incompatible API changes,
* Minor: when you add functionality in a backwards compatible manner, and
* Patch: when you make backwards compatible bug fixes.

> [https://semver.org/](https://semver.org/)

# Installation
## using go 
```
go get github.com/FalcoSuessgott/tago
```

## using curl
```
 curl -sf https://gobinaries.com/FalcoSuessgott/tago | sh
```


# Usage
```
bumping semantic versioning git tags

Usage:
  tago [flags]

Flags:
  -c, --current         print latest tag if it is standing at head
  -h, --help            help for tago
  -x, --major           bump major version part
  -y, --minor           bump minor version part
  -m, --msg string      tag annotation message
  -z, --patch           bump patch version part
      --prefix          use "v" as prefix
  -p, --push            push tag afterwards
  -r, --remote string   remote (default "origin")
```
# Examples

## bump minor version
```
tago --minor -m "added update option"
```

## bump patch version and push created tag afterwards
```
tago -pz -m "fixed authentication bug"
```
