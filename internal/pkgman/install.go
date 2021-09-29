package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/aurrpc"
	"github.com/YKMeIz/layman/internal/cmd"
	"github.com/YKMeIz/layman/internal/color"
	"github.com/go-git/go-git/v5"
	"os"
)

func init() {
	// cleanup working directory
	os.RemoveAll(os.TempDir() + "/aur/work/")
}

func (lc *LaymanConf) Install(pkgs ...string) error {
	if len(pkgs) == 0 {
		return nil
	}

	workDir := os.TempDir() + "/aur/work/"
	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		if err = os.MkdirAll(workDir, 0755); err != nil {
			return err
		}
	}

	pkgsInfo := aurrpc.Info(pkgs...)

	fmt.Println("\nFollowing packages are going to be installed:")
	for _, v := range pkgsInfo.Results {
		fmt.Println(" ", v.Name, v.Version)
	}

	if !lc.askForConfirmation("\nWould you like to install these packages?") {
		os.Exit(0)
	}

	for _, v := range pkgs {
		dir := workDir + v

		_ = os.Remove(dir)

		_, err := git.PlainClone(dir, false, &git.CloneOptions{
			URL:      "https://aur.archlinux.org/" + v,
			Progress: os.Stdout,
		})
		if err != nil {
			return err
		}

		makepkgCmd := "makepkg -sicr --noconfirm"
		if lc.Verbose {
			// --printsrcinfo will print info only, not make package
			makepkgCmd += " -L"
		}
		if lc.SkipPGPCheck {
			makepkgCmd += " --skippgpcheck"
		}

		if err := cmd.ExecCmd(dir, makepkgCmd); err != nil {
			if !lc.Force {
				return err
			}
		}

		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func (lc *LaymanConf) Remove(pkgs ...string) error {
	fmt.Println("\nFollowing packages are going to be removed:")

	if !lc.Force {
		for _, v := range pkgs {
			if info, ok := lc.Installed[v]; ok {
				fmt.Println(" ", v, info)
				continue
			}
			println(color.Red("Error: package", v, "not found in world"))
			os.Exit(-1)
		}
	}

	if !lc.askForConfirmation("\nWould you like to remove these packages?") {
		os.Exit(0)
	}

	for _, v := range pkgs {
		if err := cmd.ExecCmd("", "sudo pacman -Rs --noconfirm "+v); err != nil {
			return err
		}
	}
	return nil
}
