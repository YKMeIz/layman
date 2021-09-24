package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/aurrpc"
	"strconv"
)

func (lc *LaymanConf) Search(pkgs ...string) {
	var count int

	for _, v := range pkgs {
		res := aurrpc.Search(v)
		count += res.Resultcount
		for _, v := range res.Results {
			installed, ok := lc.Installed[v.Name]
			if !ok {
				installed = "[ Not Installed ]"
			}
			fmt.Printf("%s\n  Latest version available: %s\n  Latest version installed: %s\n  Homepage: %s\n  Description: %s\n\n",
				v.Name, v.Version, installed, v.URL, v.Description)
		}
	}

	fmt.Println("[ Applications found :", strconv.Itoa(count), "]")
}
