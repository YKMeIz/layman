package aurrpc

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func Search(pkg string) SearchRPC {
	var res SearchRPC

	url := "https://aur.archlinux.org/rpc/?v=5&type=search&arg=" + pkg

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

	return res
}
