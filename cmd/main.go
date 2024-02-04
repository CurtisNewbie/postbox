package main

import (
	"os"

	"github.com/curtisnewbie/postbox/internal/postbox"
)

func main() {
	postbox.BootstrapServer(os.Args)
}
