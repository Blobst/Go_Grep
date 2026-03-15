package internal

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"strings"
)

type CLIOptions struct {
	IgnoreCase bool
	Version    bool
	Update     bool
	Help       bool
	Pattern    string
	Path       string
}

func Usage() string {
	return strings.TrimSpace(`Usage: gogrep [flags] <pattern> [path]

Search recursively for a regular expression pattern.

Flags:
  -i            Perform case-insensitive search
  -v, --version Print version and exit
  -u, --update  Check for updates and exit
  -h, --help    Show this help message
`)
}

func ParseCLIArgs(args []string) (CLIOptions, error) {
	fs := flag.NewFlagSet("gogrep", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	ignoreCase := fs.Bool("i", false, "Perform case-insensitive search")
	versionShort := fs.Bool("v", false, "Print version and exit")
	versionLong := fs.Bool("version", false, "Print version and exit")
	updateShort := fs.Bool("u", false, "Check for updates and exit")
	updateLong := fs.Bool("update", false, "Check for updates and exit")
	helpShort := fs.Bool("h", false, "Show help")
	helpLong := fs.Bool("help", false, "Show help")

	if err := fs.Parse(args); err != nil {
		return CLIOptions{}, err
	}

	options := CLIOptions{
		IgnoreCase: *ignoreCase,
		Version:    *versionShort || *versionLong,
		Update:     *updateShort || *updateLong,
		Help:       *helpShort || *helpLong,
		Path:       ".",
	}

	remaining := fs.Args()
	if options.Help || options.Version || options.Update {
		return options, nil
	}

	if len(remaining) == 0 {
		return CLIOptions{}, errors.New("missing search pattern")
	}

	if len(remaining) > 2 {
		return CLIOptions{}, fmt.Errorf("too many arguments: %s", strings.Join(remaining[2:], " "))
	}

	options.Pattern = remaining[0]
	if len(remaining) == 2 {
		options.Path = remaining[1]
	}

	return options, nil
}
