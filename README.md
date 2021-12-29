# XGiF (Exposed Git Finder)
## About XGiF
XGiF (Exposed Git Finder) is tool written in Go designed to find .git folder exposed due to server misconfiguration. Such misconfiguration in a web application can lead to source code disclosure and invite other serious vulnerabilities.

## Screenshots
![XGiF](https://github.com/prasant-paudel/XGiF/raw/main/screenshot.png "XGiF Screenshot")

## Installation
```
go install github.com/prasant-paudel/xgif@latest
```

## Usage
Flag | Description 
-----|-------------
-t | Target URL 
-T | List of target URLs
-v | Enable verbose mode (default=false)
-o | Output to a file

## Examples
```sh
xgif -t https://example.com
```
```sh
xgif -T target_urls.txt
```
```sh
xgif -T target_urls.txt -v -o output_file.txt
```

## Version
**Current Version is 1.1**