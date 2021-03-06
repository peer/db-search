package main

import (
	"github.com/alecthomas/kong"
	"gitlab.com/tozd/go/errors"

	"gitlab.com/peerdb/search/internal/cli"
)

func main() {
	var config Config
	cli.Run(&config, "", func(_ *kong.Context) errors.E {
		return generate(&config)
	})
}
