package main

import (
	"github.com/YKMeIz/layman/internal/color"
	"github.com/YKMeIz/layman/internal/pkgman"
	"github.com/spf13/cobra"
	"os"
)

type flagSet struct {
	remove, update, list bool
	lc                   *pkgman.LaymanConf
}

func main() {
	var (
		fs flagSet

		cmd = &cobra.Command{
			Use:     "layman",
			Short:   "Layman is a package manager for Arch User Repository (AUR).",
			Version: "0.0.1",
			Run:     fs.execute,
		}
	)

	fs.lc = pkgman.New()

	cmd.PersistentFlags().BoolVarP(&fs.lc.Ask, "ask", "a", false, "before performing the action, display what will take place")
	cmd.PersistentFlags().BoolVarP(&fs.remove, "clean", "c", false, "clean the system by removing all matching packages")
	cmd.PersistentFlags().BoolVarP(&fs.update, "update", "u", false, "update packages to the latest version available")
	cmd.PersistentFlags().BoolVarP(&fs.list, "list", "l", false, "display a list of installed packages")
	cmd.PersistentFlags().BoolVarP(&fs.lc.Verbose, "verbose", "v", false, "tell layman to run in verbose mode")
	cmd.PersistentFlags().BoolVar(&fs.lc.SkipPGPCheck, "skippgpcheck", false, "do not verify PGP signatures of source files")
	cmd.PersistentFlags().BoolVar(&fs.lc.Force, "force", false, "ignore errors returned by pacman")

	if err := cmd.Execute(); err != nil {
		println(color.Red(err.Error()))
		os.Exit(1)
	}
}

func (fs *flagSet) execute(cmd *cobra.Command, args []string) {

	if fs.remove {
		if fs.update {
			println(color.Red("Error: cannot remove packages with update together"))
			os.Exit(1)
		}
		if fs.list {
			println(color.Red("Error: cannot remove packages with list together"))
			os.Exit(1)
		}

		if err := fs.lc.Remove(args...); err != nil {
			println(err.Error())
			os.Exit(1)
		}

		return
	}

	if fs.update {
		if fs.remove {
			println(color.Red("Error: cannot update packages with clean together"))
			os.Exit(1)
		}
		if fs.list {
			println(color.Red("Error: cannot update packages with list together"))
			os.Exit(1)
		}

		fs.lc.Update(args...)
		return
	}

	if fs.list {
		if fs.update {
			println(color.Red("Error: cannot list packages with update together"))
			os.Exit(1)
		}
		if fs.remove {
			println(color.Red("Error: cannot list packages with clean together"))
			os.Exit(1)
		}

		fs.lc.List()
		return
	}

	if len(args) > 0 {
		if err := fs.lc.Install(args...); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
}
