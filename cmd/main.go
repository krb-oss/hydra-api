package main

import (
	"fmt"
	"os"

	hydra_api "github.com/krb-oss/hydra-api"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	app := hydra_api.Fiber()
	srv := hydra_api.New(app)
	return srv.Listen()
}
