package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/cmd"
	"github.com/YKMeIz/layman/internal/color"
	"github.com/YKMeIz/layman/internal/config"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
)

func init() {
	// cleanup working directory
	os.RemoveAll(os.TempDir() + "/aur/work/")
}

func Install(pkgs ...string) error {
	if len(pkgs) == 0 {
		return nil
	}

	workDir := os.TempDir() + "/aur/work/"
	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		if err = os.MkdirAll(workDir, 0755); err != nil {
			return err
		}
	}

	for _, v := range pkgs {
		if _, err := getLatestVersion(v); err != nil {
			return err
		}
		fmt.Println(v)
	}

	if !askForConfirmation("Would you like to install these packages?") {
		os.Exit(0)
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

		makepkgCmd := "makepkg -sicr --noconfirm"
		if verboseMode {
			makepkgCmd += " -L --printsrcinfo"
		}
		if skippgpcheck {
			makepkgCmd += " --skippgpcheck"
		}

		if err := cmd.ExecCmd(dir, makepkgCmd); err != nil {
			if !force {
				return err
			}
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
	if !force {
		for _, v := range pkgs {
			if !config.IsExist(v) {
				println(color.Red("Error: package", v, "not found in world"))
				os.Exit(-1)
			}
			fmt.Println(v)
		}
	}

	if !askForConfirmation("Would you like to remove these packages?") {
		os.Exit(0)
	}

	for _, v := range pkgs {
		if force {
			config.RemovePackage(v)
		}
		if err := cmd.ExecCmd("", "sudo pacman -Rs --noconfirm "+v); err != nil {
			return err
		}

		if !force {
			config.RemovePackage(v)
		}
	}
	return nil
}
