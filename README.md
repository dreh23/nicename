# Nicename 🐶🧀

[![CI](https://github.com/dreh23/nicename/actions/workflows/ci.yml/badge.svg)](https://github.com/dreh23/nicename/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dreh23/nicename.svg)](https://pkg.go.dev/github.com/dreh23/nicename)
[![Go Report Card](https://goreportcard.com/badge/github.com/dreh23/nicename)](https://goreportcard.com/report/github.com/dreh23/nicename)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight, concurrency-safe Go library and CLI for generating delightful, human-friendly random names, URL slugs, and unique task identifiers.

Now refreshed with full support for **Wallace & Gromit** and the **Aardman Animations** filmography (*Chicken Run*, *Shaun the Sheep*, *Morph*, *Creature Comforts*, *Flushed Away*, *The Pirates!*, and more).

---

## ✨ Features

- **Zero External Dependencies**: Pure Go standard library (using modern `math/rand/v2` and `crypto/rand`).
- **High Performance & Thread-Safe**: Sub-microsecond execution, zero regex overhead, and fully concurrency-safe for high-throughput goroutines and servers.
- **Aardman Universe Support**: Over 100+ iconic characters, inventions, cheeses, and British adjectives:
  - *Characters*: Gromit, Wallace, Feathers McGraw, Wendolene Ramsbottom, Shaun the Sheep, Preston, Rocky Rhodes, Ginger, Morph, Chas, Bitzer, The Pirate Captain...
  - *Inventions & Cheeses*: Techno Trousers, Knit-o-matic, Snoozatron, Bun-vac 6000, Wensleydale, Stinking Bishop, Cracking Toast...
  - *Adjectives*: Cracking, Smashing, Plucky, Cheesy, Grand, Gallant, Clever, Feathered, Eccentric...
- **Kebab-Case Slugs & Task IDs**: Ready-made for Docker container names, Git worktree branches, and background jobs.
- **CLI Utility Included**: Generate names directly from the command line or shell scripts.
- **100% Backward Compatible**: Retains the classic `First` and `Second` name lists and `GeneratePair()`.

---

## 📦 Installation

### As a Go Library
```bash
go get github.com/dreh23/nicename
```

Requires Go 1.22 or newer.

### As a CLI Tool
```bash
go install github.com/dreh23/nicename/cmd/nicename@latest
```

---

## 💻 CLI Usage

```bash
# Default: Spirited Aardman pair
$ nicename
Cracking Gromit

# URL- and Git-branch-friendly kebab-case slug
$ nicename -slug
cracking-gromit

# Unique collision-resistant task ID (slug + 4-hex random suffix)
$ nicename -task-id
cracking-gromit-8f2a

# Specific Aardman character pair
$ nicename -character
Clever Feathers McGraw

# Classic original dataset
$ nicename -classic
Adventurous Mary

# Generate multiple names
$ nicename -count 3 -slug
plucky-feathers-mcgraw
cheesy-techno-trousers
grand-wendolene-ramsbottom
```

---

## 🚀 Quick Start (Go Code)

```go
package main

import (
	"fmt"
	"github.com/dreh23/nicename"
)

func main() {
	// 1. Classic generation (Backward compatible)
	fmt.Println(nicename.GeneratePair())
	// Output: "Adventurous Mary"

	// 2. Aardman / Wallace & Gromit generation
	fmt.Println(nicename.GenerateAardman())
	// Output: "Cracking Gromit"
	// Output: "Cheesy Techno Trousers"

	// 3. Specific Aardman character pair
	fmt.Println(nicename.GenerateAardmanCharacter())
	// Output: "Clever Feathers McGraw"

	// 4. URL and Git branch slugs
	fmt.Println(nicename.GenerateAardmanSlug())
	// Output: "cracking-gromit"
	// Output: "plucky-shaun-the-sheep"

	// 5. Unique, collision-resistant task IDs
	fmt.Println(nicename.GenerateTaskID())
	// Output: "cracking-gromit-8f2a"
	// Output: "feathers-mcgraw-4e1b"
}
```

---

## 🛠️ API Reference

| Function | Returns | Example | Description |
| :--- | :--- | :--- | :--- |
| `GeneratePair()` | `string` | `"Adventurous Mary"` | Classic adjective + name from the original 4,000+ name dictionary. |
| `GeneratePairSlug()` | `string` | `"adventurous-mary"` | Lowercase kebab-cased slug from the original dictionary. |
| `GenerateAardman()` | `string` | `"Cracking Gromit"` | Spirited Aardman pair (character, invention, or cheese). |
| `GenerateAardmanCharacter()`| `string` | `"Clever Feathers McGraw"`| Adjective paired strictly with an Aardman character. |
| `GenerateAardmanSlug()` | `string` | `"cracking-gromit"` | URL- and branch-safe kebab-cased Aardman slug. |
| `GenerateTaskID()` | `string` | `"cracking-gromit-8f2a"` | Aardman slug with a 4-hex crypto-random suffix. |
| `FormatTaskID(slug string)` | `string` | `"custom-task-8f2a"` | Appends a 4-hex crypto-random suffix to any slug. |
| `Slugify(s string)` | `string` | `"feathers-mcgraw"` | Cleans and normalizes any string into a lowercase slug. |
| `GenerateCustom(...)` | `string` | `"Custom Pair"` | Generates pairs from arbitrary custom word slices. |

---

## ⚡ Performance

Benchmarks executed on Intel Core i5-12400 (Linux amd64, pure Go):

```text
BenchmarkGeneratePair-12          22,729,735     51.2 ns/op     18 B/op    1 allocs/op
BenchmarkGenerateAardmanSlug-12   10,013,353    117.2 ns/op     47 B/op    2 allocs/op
BenchmarkGenerateTaskID-12         2,188,717    546.7 ns/op     83 B/op    5 allocs/op
```

---

## 🧪 Testing

```bash
go test -v -race ./...
```

Statement coverage: **100.0%**.

---

## 📜 License

MIT License. See [LICENSE](LICENSE) for details.
