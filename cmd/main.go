package main

import (
	"github.com/YKMeIz/layman/internal/color"
	"github.com/YKMeIz/layman/internal/pkgman"
	"github.com/spf13/cobra"
	"os"
)

type flagSet struct {
	ask, remove, update, list, verbose, skippgpcheck bool
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

	cmd.PersistentFlags().BoolVarP(&fs.ask, "ask", "a", false, "before performing the action, display what will take place")
	cmd.PersistentFlags().BoolVarP(&fs.remove, "clean", "c", false, "cleans the system by removing all matching packages")
	cmd.PersistentFlags().BoolVarP(&fs.update, "update", "u", false, "updates packages to the latest version available")
	cmd.PersistentFlags().BoolVarP(&fs.list, "list", "l", false, "displays a list of installed packages")
	cmd.PersistentFlags().BoolVarP(&fs.verbose, "verbose", "v", false, "tell layman to run in verbose mode")
	cmd.PersistentFlags().BoolVar(&fs.skippgpcheck, "skippgpcheck", false, "do not verify PGP signatures of source files")

	if err := cmd.Execute(); err != nil {
		println(color.Red(err.Error()))
		os.Exit(1)
	}
}

func (fs *flagSet) execute(cmd *cobra.Command, args []string) {
	if fs.ask {
		pkgman.SetAskMode()
	}
	if fs.verbose {
		pkgman.SetVerboseMode()
	}

	if fs.skippgpcheck {
		pkgman.SetSkipPGPCheck()
	}

	if fs.remove {
		if fs.update {
			println(color.Red("Error: cannot remove packages with update together"))
			os.Exit(1)
		}
		if fs.list {
			println(color.Red("Error: cannot remove packages with list together"))
			os.Exit(1)
		}

		if err := pkgman.Remove(args...); err != nil {
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

		pkgman.Update(args...)
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

		pkgman.List()
		return
	}

	if len(args) > 0 {
		if err := pkgman.Install(args...); err != nil {
			println(err.Error())
			os.Exit(1)
		}
	}
}
