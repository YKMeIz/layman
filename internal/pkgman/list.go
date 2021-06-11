package pkgman

import (
	"fmt"
	"github.com/YKMeIz/layman/internal/config"
)

func List() {
	config.WalkThroughKVs(func(k, v string) error {
		fmt.Println(k, v)
		return nil
	})
}
