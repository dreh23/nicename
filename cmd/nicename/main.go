package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/dreh23/nicename"
)

var osExit = os.Exit

func main() {
	osExit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("nicename", flag.ContinueOnError)
	fs.SetOutput(stderr)

	var (
		slugFlag      = fs.Bool("slug", false, "Output lowercase kebab-cased slug (e.g. cracking-gromit)")
		taskIDFlag    = fs.Bool("task-id", false, "Output collision-resistant task ID (e.g. cracking-gromit-8f2a)")
		classicFlag   = fs.Bool("classic", false, "Use classic adjective + name dataset instead of Aardman")
		characterFlag = fs.Bool("character", false, "Use strictly Aardman character names")
		countFlag     = fs.Int("count", 1, "Number of names to generate (must be non-negative)")
	)

	fs.Usage = func() {
		fmt.Fprintf(stderr, "Usage: nicename [options]\n\n")
		fmt.Fprintf(stderr, "A lightweight generator for memorable names, slugs, and task IDs.\n\n")
		fmt.Fprintf(stderr, "Options:\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		if err == flag.ErrHelp {
			return 0
		}
		return 2
	}

	if *countFlag < 0 {
		fmt.Fprintf(stderr, "Error: -count must be non-negative, got %d\n", *countFlag)
		return 1
	}

	if *classicFlag && *characterFlag {
		fmt.Fprintf(stderr, "Error: cannot combine -classic and -character flags\n")
		return 1
	}

	out := bufio.NewWriter(stdout)
	defer out.Flush()

	for i := 0; i < *countFlag; i++ {
		var base string
		switch {
		case *classicFlag:
			base = nicename.GeneratePair()
		case *characterFlag:
			base = nicename.GenerateAardmanCharacter()
		default:
			base = nicename.GenerateAardman()
		}

		var line string
		switch {
		case *taskIDFlag:
			line = nicename.FormatTaskID(nicename.Slugify(base))
		case *slugFlag:
			line = nicename.Slugify(base)
		default:
			line = base
		}

		fmt.Fprintln(out, line)
	}

	return 0
}
