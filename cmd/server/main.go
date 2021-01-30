package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/seandrucker/golang-starter-cmd/pkg/app"
)

func main() {
	// set usage first, emitted on any -h or an error in flag.Parse below
	flag.CommandLine.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println()
		fmt.Println(" ", filepath.Base(os.Args[0]), "[-v][-l address]")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println()
		flag.PrintDefaults()
	}

	// not always this fortunate to parse directly from command line into app struct.
	app := &app.App{}
	flag.BoolVar(&app.Verbose, "v", false, "turn on verbose output")
	flag.StringVar(&app.Addr, "l", ":8080", "listen on network address")
	flag.Parse()

	// run app
	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
