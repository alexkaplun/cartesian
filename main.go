package main

import (
	"github.com/alexkaplun/cartesian/cli"
	"os"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
