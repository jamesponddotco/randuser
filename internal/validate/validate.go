// Package validate contains a single function to validate user input.
package validate

import (
	"slices"
	"strings"
)

// Format returns false if the given input doesn't comply with the format
// expected by the application.
func Format(input string) bool {
	input = strings.ReplaceAll(input, " ", "")

	if !strings.Contains(input, "{{") && !strings.Contains(input, "}}") {
		return false
	}

	var (
		placeholders = []string{
			"adjective",
			"noun",
			"verb",
			"number",
		}
		segments = strings.Split(input, "{{")
	)

	for i := 1; i < len(segments); i++ {
		parts := strings.Split(segments[i], "}}")

		if len(parts) != 2 {
			return false
		}

		placeholder := parts[0]
		if !slices.Contains(placeholders, placeholder) {
			return false
		}
	}

	return true
}
