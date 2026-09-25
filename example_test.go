package nicename_test

import (
	"fmt"
	"strings"

	"github.com/dreh23/nicename"
)

func ExampleGeneratePair() {
	pair := nicename.GeneratePair()
	// Output will be something like "Adventurous Mary"
	fmt.Println(len(strings.Split(pair, " ")) >= 2)
	// Output:
	// true
}

func ExampleGenerateAardman() {
	aardman := nicename.GenerateAardman()
	// Output will be something like "Cracking Gromit"
	fmt.Println(aardman != "")
	// Output:
	// true
}

func ExampleGenerateAardmanSlug() {
	slug := nicename.GenerateAardmanSlug()
	// Output will be something like "cracking-gromit"
	fmt.Println(strings.Contains(slug, "-"))
	// Output:
	// true
}

func ExampleGenerateTaskID() {
	taskID := nicename.GenerateTaskID()
	// Output will be something like "cracking-gromit-8f2a"
	fmt.Println(len(taskID) > 5)
	// Output:
	// true
}

func ExampleFormatTaskID() {
	taskID := nicename.FormatTaskID("custom-worker-task")
	// Output will be something like "custom-worker-task-8f2a"
	fmt.Println(strings.HasPrefix(taskID, "custom-worker-task-"))
	// Output:
	// true
}

func ExampleGenerateCustom() {
	adjectives := []string{"Grand", "Smashing"}
	nouns := []string{"Contraption", "Rocket"}

	name := nicename.GenerateCustom(adjectives, nouns, "-", true)
	// Returns lowercase slug joined by hyphen
	fmt.Println(strings.HasPrefix(name, "grand-") || strings.HasPrefix(name, "smashing-"))
	// Output:
	// true
}

func ExampleSlugify() {
	fmt.Println(nicename.Slugify("Feathers McGraw"))
	fmt.Println(nicename.Slugify("Bun-vac 6000"))
	fmt.Println(nicename.Slugify("  Cracking   Toast!  "))
	// Output:
	// feathers-mcgraw
	// bun-vac-6000
	// cracking-toast
}
