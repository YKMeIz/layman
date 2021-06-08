package pkgman

import (
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
)

func getLatestVersion(pkg string) (string, error) {
	r := git.NewRemote(memory.NewStorage(), &config.RemoteConfig{
		URLs: []string{"https://aur.archlinux.org/" + pkg + ".git"},
	})

	rfs, err := r.List(&git.ListOptions{})
	if err != nil {
		return "", err
	}

	for _, v := range rfs {
		if v.Name() == plumbing.Master {
			return v.Hash().String(), nil
		}
	}

	return "", errors.New("cannot find any version information for package: " + pkg)
}
