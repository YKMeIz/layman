package pkgman

import (
	"github.com/YKMeIz/layman/internal/cmd"
	"os"
	"strings"
)

const (
	listAURPkgCmd = "pacman -Qem"
)

func (lc *LaymanConf) List() {
	if err := cmd.ExecCmd("", listAURPkgCmd); err != nil {
		os.Exit(1)
	}
}

func retrievePkgList() map[string]string {
	res := make(map[string]string)

	b, err := cmd.ExecCmdOutput("", listAURPkgCmd)
	if err != nil {
		panic(err)
	}

	for _, v := range strings.Split(string(b), "\n") {
		s := strings.Split(v, " ")
		// each entity is in format, e.g. "typora 0.11.8-1"
		if len(s) != 2 {
			continue
		}

		res[s[0]] = s[1]
	}

	return res
}
