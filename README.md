# XGiF (Exposed Git Finder)
## About XGiF
XGiF (Exposed Git Finder) is tool written in Go designed to find .git folder exposed due to server misconfiguration. Such misconfiguration in a web application can lead to source code disclosure and invite other serious vulnerabilities.

## Screenshots
![XGiF](https://github.com/prasant-paudel/XGiF/raw/main/screenshot.png "XGiF Screenshot")

## Installation
```
go get -u github.com/prasant-paudel/xgif
```

## Usage
```sh
xgif [FILE]
```
## Example
```sh
xgif target_urls.txt
```

## Version
**Current Version is 1.0**
