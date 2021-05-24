# XGiF (Exposed Git Finder)
## About XGiF
XGiF (Exposed Git Finder) is a python tool designed to find .git folder exposed due to server misconfiguration. Such misconfiguration in a web application can lead to source code disclosure.

It check for the following files and folders:
- `/.git/`
- `/.git/config`
- `/.git/HEAD`
- `/.git/logs/HEAD`
- `/.git/index`

## Installation
```
git clone https://github.com/prasant-paudel/XGiF.git
```

## Dependencies
XGiF depends upon the `request`, `argparse` and `colorama` python modules.

These dependencies can be installed using the requirements file:

- Installation on Windows:
```
c:\python27\python.exe -m pip install -r requirements.txt
```
- Installation on Linux
```
sudo pip install -r requirements.txt
```
## Usage
Short  from | Long form | Description
------------|-----------|------------
-d | --domain | Check a single domain
-D | --domains | Check domains from file
-t | --threads | Number of threads to use (default=40)
-v | --verbose | Enable Verbosity
-h | --help | Show the help message and exit.

## Examples

Single Domain:
- `python xgif.py -d example.com`
or 
- `python xgif.py -d http://example.com/`

Multiple domains from a file:
- `python xgif.py -D domains.txt`

Define custom number of threads to use:
- `python xgif.py -D domains.txt -t 40`

## Version
**Current Version is 1.0**
