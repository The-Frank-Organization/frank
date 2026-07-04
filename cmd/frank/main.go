package main

import (
	"flag"
	"fmt"
	"os"

	frankrecover "github.com/jackli/frank/internal/recover"
)

func main() {
	root := flag.String("store", ".frank-store", "store root")
	flag.Parse()
	if err := frankrecover.Run(*root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
