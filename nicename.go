// Package nicename provides lightweight, thread-safe generators for
// memorable, human-friendly random names, URL slugs, and unique task identifiers.
package nicename

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	randv2 "math/rand/v2"
	"strings"
)

// GeneratePair generates a classic random adjective + name pair
// from the original dataset (e.g. "Adventurous Mary").
// Preserves full backwards compatibility.
func GeneratePair() string {
	if len(First) == 0 || len(Second) == 0 {
		return ""
	}
	first := First[randv2.IntN(len(First))]
	second := Second[randv2.IntN(len(Second))]
	return first + " " + second
}

// GeneratePairSlug generates a lowercase, kebab-cased slug from the original dataset
// (e.g. "adventurous-mary").
func GeneratePairSlug() string {
	return Slugify(GeneratePair())
}

// GenerateAardman generates a spirited Aardman / Wallace & Gromit themed pair
// (e.g. "Cracking Gromit", "Cheesy Techno Trousers").
func GenerateAardman() string {
	if len(AardmanAdjectives) == 0 || len(AardmanNouns) == 0 {
		return ""
	}
	adj := AardmanAdjectives[randv2.IntN(len(AardmanAdjectives))]
	noun := AardmanNouns[randv2.IntN(len(AardmanNouns))]
	return adj + " " + noun
}

// GenerateAardmanCharacter generates a pair pairing an adjective with a specific
// character from the Aardman universe (e.g. "Clever Feathers McGraw").
func GenerateAardmanCharacter() string {
	if len(AardmanAdjectives) == 0 || len(AardmanCharacters) == 0 {
		return ""
	}
	adj := AardmanAdjectives[randv2.IntN(len(AardmanAdjectives))]
	char := AardmanCharacters[randv2.IntN(len(AardmanCharacters))]
	return adj + " " + char
}

// GenerateAardmanSlug generates a URL- and Git-branch-friendly slug from
// the Aardman dataset (e.g. "cracking-gromit", "plucky-feathers-mcgraw").
func GenerateAardmanSlug() string {
	return Slugify(GenerateAardman())
}

// GenerateTaskID returns a unique, collision-resistant identifier ideal for
// worker tasks, git branches, and container names (e.g. "cracking-gromit-8f2a").
func GenerateTaskID() string {
	slug := GenerateAardmanSlug()
	var suffixBytes [2]byte
	if _, err := crand.Read(suffixBytes[:]); err != nil {
		// Fallback to PRNG if crypto/rand is unavailable
		return fmt.Sprintf("%s-%04x", slug, randv2.Uint32()&0xffff)
	}
	return slug + "-" + hex.EncodeToString(suffixBytes[:])
}

// GenerateCustom picks an adjective and a noun from custom slices, joining them with separator.
func GenerateCustom(adjectives []string, nouns []string, separator string, toSlug bool) string {
	if len(adjectives) == 0 || len(nouns) == 0 {
		return ""
	}
	adj := adjectives[randv2.IntN(len(adjectives))]
	noun := nouns[randv2.IntN(len(nouns))]
	res := adj + separator + noun
	if toSlug {
		return Slugify(res)
	}
	return res
}

// Slugify converts any string into a clean lowercase, hyphen-separated slug,
// stripping special punctuation while preserving alphanumeric words.
func Slugify(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	inHyphen := false

	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'):
			b.WriteByte(c)
			inHyphen = false
		case c >= 'A' && c <= 'Z':
			b.WriteByte(c + ('a' - 'A'))
			inHyphen = false
		default:
			if !inHyphen && b.Len() > 0 {
				b.WriteByte('-')
				inHyphen = true
			}
		}
	}
	return strings.TrimSuffix(b.String(), "-")
}
