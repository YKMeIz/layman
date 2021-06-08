package pkgman

import (
	"github.com/YKMeIz/layman/internal/config"
	"os"
)

func Update() {
	var outDatedPkgs []string

	config.UpdatePackage(func(name, version string) string {
		newVersion, err := getLatestVersion(name)
		if err != nil {
			println(err.Error())
			return ""
		}

		if newVersion != version {
			outDatedPkgs = append(outDatedPkgs, name)
		}

		return ""
	})

	if err := Install(outDatedPkgs...); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}
