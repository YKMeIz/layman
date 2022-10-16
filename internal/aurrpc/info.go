package aurrpc

import (
	"encoding/json"
	"fmt"
	"github.com/YKMeIz/layman/internal/color"
	"io"
	"log"
	"net/http"
	"os"
)

func Info(pkgs ...string) InfoRPC {
	var (
		res         InfoRPC
		notComplete bool
	)

	url := "https://aur.archlinux.org/rpc/?v=5&type=info"
	for _, v := range pkgs {
		url += "&arg[]=" + urlNormalization(v)[0]
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range pkgs {
		if !search(v, res) {
			fmt.Println(color.Red("cannot find package ", v))
			notComplete = true
		}
	}

	if notComplete {
		os.Exit(1)
	}

	return res
}

func search(name string, info InfoRPC) bool {
	for _, v := range info.Results {
		if v.Name == name {
			return true
		}
	}

	return false
}
