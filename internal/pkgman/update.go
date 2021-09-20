package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/aurrpc"
	"github.com/YKMeIz/layman/internal/color"
	"os"
)

func (lc *LaymanConf) Update(pkgs ...string) {
	outDatedPkgs := lc.findOutdated(pkgs...)

	if err := lc.Install(outDatedPkgs...); err != nil {
		println(err.Error())
		os.Exit(1)
	}
}

func (lc *LaymanConf) findOutdated(pkgs ...string) []string {
	var (
		info aurrpc.InfoRPC
		res  []string
	)

	if len(pkgs) == 0 {
		for k, _ := range lc.Installed {
			pkgs = append(pkgs, k)
		}
	}

	info = aurrpc.Info(pkgs...)
	for _, v := range info.Results {

		if k, ok := lc.Installed[v.Name]; ok && k == v.Version {
			if lc.Verbose {
				fmt.Println("package", color.Bold(v.Name), "is on the latest version")
			}
			continue
		}

		res = append(res, v.Name)
		if lc.Verbose {
			if k, ok := lc.Installed[v.Name]; ok {
				fmt.Println("find latest version for package "+color.Bold(v.Name)+":", color.Blue(k), "->", color.Yellow(v.Version))
				continue
			}
			fmt.Println("find latest version for package "+color.Bold(v.Name)+": ->", color.Yellow(v.Version))
		}
	}

	return res
}
