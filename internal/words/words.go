// Package words contains embedded word lists.
package words

import (
	_ "embed"
	"strings"
)

//go:embed data/adjectives.txt
var _adjectives string

//go:embed data/verbs.txt
var _verbs string

//go:embed data/nouns.txt
var _nouns string

// Adjectives returns a list of adjectives.
func Adjectives() []string {
	return strings.Split(_adjectives, "\n")
}

// Verbs returns a list of verbs.
func Verbs() []string {
	return strings.Split(_verbs, "\n")
}

// Nouns returns a list of nouns.
func Nouns() []string {
	return strings.Split(_nouns, "\n")
}
