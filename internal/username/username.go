// Package username contains a single function to generate random usernames.
package username

import (
	"strings"

	"git.sr.ht/~jamesponddotco/acopw-go"
	"git.sr.ht/~jamesponddotco/randuser/internal/words"
	"git.sr.ht/~jamesponddotco/xstd-go/xerrors"
	"git.sr.ht/~jamesponddotco/xstd-go/xstrings"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ErrInvalidFormat is returned when an invalid format is provided.
const ErrInvalidFormat xerrors.Error = "invalid format; must be a string with the following placeholders: {{adjective}}, {{verb}}, {{noun}}, {{number}}"

// Generate generates a random username based on the given format without title
// case.
func Generate(format string) (string, error) {
	generated, err := generateWords(format)
	if err != nil {
		return "", err
	}

	return xstrings.Join(generated...), nil
}

// GenerateWithTitleCase generates a random username based on the given format
// with title case.
func GenerateWithTitleCase(format string) (string, error) {
	generated, err := generateWords(format)
	if err != nil {
		return "", err
	}

	titleCaser := cases.Title(language.English)
	for i, word := range generated {
		generated[i] = titleCaser.String(word)
	}

	return xstrings.Join(generated...), nil
}

func generateWords(format string) ([]string, error) {
	format = strings.ReplaceAll(format, " ", "")

	var neededGenerators []string

	for _, placeholder := range []string{
		"{{adjective}}",
		"{{verb}}",
		"{{noun}}",
		"{{number}}",
	} {
		if strings.Contains(format, placeholder) {
			neededGenerators = append(neededGenerators, placeholder)
		}
	}

	if len(neededGenerators) == 0 {
		return nil, ErrInvalidFormat
	}

	generators := map[string]func() string{
		"{{adjective}}": func() string {
			generator := &acopw.Diceware{
				Length:     1,
				Capitalize: false,
				Words:      words.Adjectives(),
			}

			return generator.Generate()
		},
		"{{verb}}": func() string {
			generator := &acopw.Diceware{
				Length:     1,
				Capitalize: false,
				Words:      words.Verbs(),
			}

			return generator.Generate()
		},
		"{{noun}}": func() string {
			generator := &acopw.Diceware{
				Length:     1,
				Capitalize: false,
				Words:      words.Nouns(),
			}

			return generator.Generate()
		},
		"{{number}}": func() string {
			generator := &acopw.PIN{
				Length: 4,
			}

			return generator.Generate()
		},
	}

	generatedWords := make([]string, 0, len(neededGenerators))

	for _, placeholder := range neededGenerators {
		word := generators[placeholder]()
		generatedWords = append(generatedWords, word)
	}

	return generatedWords, nil
}
