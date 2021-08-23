# layman

Layman is a package manager for Arch User Repository (AUR).

## Usage
```
Usage:
  layman [flags]

Flags:
  -a, --ask            before performing the action, display what will take place
  -c, --clean          clean the system by removing all matching packages
      --force          ignore errors returned by pacman
  -h, --help           help for layman
  -l, --list           display a list of installed packages
      --skippgpcheck   do not verify PGP signatures of source files
  -u, --update         update packages to the latest version available
  -v, --verbose        tell layman to run in verbose mode
      --version        version for layman
```

## Build

To build layman, simply us `go`:
```
$ go build -o layman cmd/main.go
```
