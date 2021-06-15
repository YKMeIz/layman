# layman

Layman is a package manager for Arch User Repository (AUR).

## Usage
```
Usage:
  layman [flags]

Flags:
  -a, --ask       before performing the action, display what will take place
  -c, --clean     cleans the system by removing all matching packages
  -h, --help      help for layman
  -l, --list      displays a list of installed packages
  -u, --update    updates packages to the latest version available
  -v, --verbose   tell layman to run in verbose mode

```

## Build

To build layman, simply us `go`:
```
$ CGO_ENABLED=0 go build -o layman cmd/main.go
```
