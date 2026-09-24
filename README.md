# Nicename 🐶🧀

[![Go Reference](https://pkg.go.dev/badge/github.com/dreh23/nicename.svg)](https://pkg.go.dev/github.com/dreh23/nicename)
[![Go Report Card](https://goreportcard.com/badge/github.com/dreh23/nicename)](https://goreportcard.com/report/github.com/dreh23/nicename)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

A lightweight, concurrency-safe Go library for generating delightful, human-friendly random names, URL slugs, and unique task identifiers.

Now refreshed with full support for **Wallace & Gromit** and the **Aardman Animations** filmography (*Chicken Run*, *Shaun the Sheep*, *Morph*, *Creature Comforts*, *Flushed Away*, *The Pirates!*, and more).

---

## ✨ Features

- **Zero External Dependencies**: Pure Go standard library (using modern `math/rand/v2` and `crypto/rand`).
- **Thread-Safe**: Fully concurrency-safe for high-throughput goroutines and servers.
- **Aardman Universe Support**: Over 100+ iconic characters, inventions, cheeses, and British adjectives:
  - *Characters*: Gromit, Wallace, Feathers McGraw, Wendolene Ramsbottom, Shaun the Sheep, Preston, Rocky Rhodes, Ginger, Morph, Chas, Bitzer, The Pirate Captain...
  - *Inventions & Cheeses*: Techno Trousers, Knit-o-matic, Snoozatron, Bun-vac 6000, Wensleydale, Stinking Bishop, Cracking Toast...
  - *Adjectives*: Cracking, Smashing, Plucky, Cheesy, Grand, Gallant, Clever, Feathered, Eccentric...
- **Kebab-Case Slugs & Task IDs**: Ready-made for Docker container names, Git worktree branches, and background jobs.
- **100% Backward Compatible**: Retains the classic `First` and `Second` name lists and `GeneratePair()`.

---

## 📦 Installation

```bash
go get github.com/dreh23/nicename
```

Requires Go 1.22 or newer.

---

## 🚀 Quick Start

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
| `Slugify(s string)` | `string` | `"feathers-mcgraw"` | Cleans and normalizes any string into a lowercase slug. |
| `GenerateCustom(...)` | `string` | `"Custom Pair"` | Generates pairs from arbitrary custom word slices. |

---

## ⚡ Performance

Benchmarks executed on Intel Core i5-12400 (Linux amd64):

```text
BenchmarkGeneratePair-12          6,695,198   173.2 ns/op
BenchmarkGenerateAardmanSlug-12      73,599  16,579 ns/op
BenchmarkGenerateTaskID-12           64,429  18,226 ns/op
```

---

## 🧪 Testing

```bash
go test -v -race ./...
```

Statement coverage: **97.0%**.

---

## 📜 License

MIT License. See [LICENSE](LICENSE) for details.
