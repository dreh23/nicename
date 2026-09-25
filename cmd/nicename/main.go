package main

import (
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
		countFlag     = flag.Int("count", 1, "Number of names to generate")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nicename [options]\n\n")
		fmt.Fprintf(os.Stderr, "A lightweight generator for memorable names, slugs, and task IDs.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *countFlag < 1 {
		*countFlag = 1
	}

	for i := 0; i < *countFlag; i++ {
		var output string
		switch {
		case *taskIDFlag:
			output = nicename.GenerateTaskID()
		case *classicFlag && *slugFlag:
			output = nicename.GeneratePairSlug()
		case *classicFlag:
			output = nicename.GeneratePair()
		case *characterFlag && *slugFlag:
			output = nicename.Slugify(nicename.GenerateAardmanCharacter())
		case *characterFlag:
			output = nicename.GenerateAardmanCharacter()
		case *slugFlag:
			output = nicename.GenerateAardmanSlug()
		default:
			output = nicename.GenerateAardman()
		}
		fmt.Println(output)
	}
}
