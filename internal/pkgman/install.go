package pkgman

import (
	"github.com/YKMeIz/layman/internal/cmd"
	"github.com/YKMeIz/layman/internal/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

func Install(pkgs ...string) error {
	workDir := os.TempDir() + "/aur/work/"
	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		if err = os.MkdirAll(workDir, 0755); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}

	for _, v := range pkgs {
		dir := workDir + v

		_ = os.Remove(dir)

		repo, err := git.PlainClone(dir, false, &git.CloneOptions{
			URL:      "https://aur.archlinux.org/" + v,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}
		ref, err := repo.Reference(plumbing.Master, true)
		if err != nil {
			return err
		}

		if err := cmd.ExecCmd(dir, "makepkg -sicr --noconfirm"); err != nil {
			return err
		}

		config.AddPackage(config.PkgInfo{
			Name:      v,
			Installed: ref.Hash().String(),
		})

		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func Remove(pkgs ...string) error {
	for _, v := range pkgs {
		if err := cmd.ExecCmd("", "sudo pacman -Rs --noconfirm "+v); err != nil {
			return err
		}

		config.RemovePackage(v)
	}
	return nil
}