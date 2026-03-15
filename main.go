package main

import (
	"fmt"
	ut "gogrep/internal"
	"os"
)

func main() {
	options, err := ut.ParseCLIArgs(os.Args[1:])
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		_, _ = fmt.Fprintln(os.Stderr, ut.Usage())
		os.Exit(2)
	}

	if options.Help {
		fmt.Println(ut.Usage())
		return
	}

	if options.Version {
		fmt.Printf("go_grep %s\n", ut.Version)
		return
	}

	if options.Update {
		ut.CheckForUpdates()
		return
	}

	found, err := ut.Search(options.Pattern, options.Path, options.IgnoreCase)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if !found {
		os.Exit(1)
	}
}
