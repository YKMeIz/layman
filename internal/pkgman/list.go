package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/color"
	"github.com/YKMeIz/layman/internal/config"
)

func List() {
	config.WalkThroughKVs(func(k, v string) error {
		fmt.Println(color.Bold(k), color.Green(v[:7]))
		return nil
	})
}
