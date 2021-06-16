package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/config"
	"os"
)

func Update(pkgs ...string) {
	if len(pkgs) == 0 {
		updateAll()
		return
	}

	var outDatedPkgs []string

	for _, v := range pkgs {
		newVersion, err := getLatestVersion(v)
		if err != nil {
			println(err.Error())
			os.Exit(1)
		}

		if verboseMode {
			fmt.Println("check latest version for package", v)
		}

		ver := config.GetVersion(v)

		if newVersion != ver {
			outDatedPkgs = append(outDatedPkgs, v)
			if verboseMode {
				fmt.Println("find latest version for package "+v+":", ver, "->", newVersion)
			}
		}

		if verboseMode {
			fmt.Println("package", v, "is on the latest version")
		}
	}

	if err := Install(outDatedPkgs...); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func updateAll() {
	var outDatedPkgs []string

	config.UpdatePackage(func(name, version string) string {
		if verboseMode {
			fmt.Println("check latest version for package", name)
		}

		newVersion, err := getLatestVersion(name)
		if err != nil {
			println(err.Error())
			return ""
		}

		if newVersion != version {
			outDatedPkgs = append(outDatedPkgs, name)
			if verboseMode {
				fmt.Println("find latest version for package "+name+":", version, "->", newVersion)
			}
		}

		if verboseMode {
			fmt.Println("package", name, "is on the latest version")
		}

		return ""
	})

	if err := Install(outDatedPkgs...); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
