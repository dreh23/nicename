package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/dreh23/nicename"
)

func main() {
	var (
		slugFlag      = flag.Bool("slug", false, "Output lowercase kebab-cased slug (e.g. cracking-gromit)")
		taskIDFlag    = flag.Bool("task-id", false, "Output collision-resistant task ID (e.g. cracking-gromit-8f2a)")
		classicFlag   = flag.Bool("classic", false, "Use classic adjective + name dataset instead of Aardman")
		characterFlag = flag.Bool("character", false, "Use strictly Aardman character names")
		countFlag     = flag.Int("count", 1, "Number of names to generate (must be non-negative)")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nicename [options]\n\n")
		fmt.Fprintf(os.Stderr, "A lightweight generator for memorable names, slugs, and task IDs.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *countFlag < 0 {
		fmt.Fprintf(os.Stderr, "Error: -count must be non-negative, got %d\n", *countFlag)
		os.Exit(1)
	}

	if *classicFlag && *characterFlag {
		fmt.Fprintf(os.Stderr, "Error: cannot combine -classic and -character flags\n")
		os.Exit(1)
	}

	out := bufio.NewWriter(os.Stdout)
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
}
