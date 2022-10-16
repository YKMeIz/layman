package aurrpc

import "strings"

func urlNormalization(pkgs ...string) []string {
	for k, v := range pkgs {
		if strings.Contains(v, "+") {
			pkgs[k] = strings.ReplaceAll(v, "+", "%2b")
		}
	}
	return pkgs
}
