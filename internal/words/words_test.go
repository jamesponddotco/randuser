package words_test

import (
	"slices"
	"testing"

	"git.sr.ht/~jamesponddotco/randuser/internal/words"
)

func TestAdjectives(t *testing.T) {
	t.Parallel()

	t.Run("returns a list of adjectives", func(t *testing.T) {
		t.Parallel()

		got := words.Adjectives()

		if len(got) == 0 {
			t.Error("expected a list of adjectives")
		}

		if !slices.Contains(got, "affected") {
			t.Error("expected the list to include the word 'affected'")
		}
	})
}

func TestVerbs(t *testing.T) {
	t.Parallel()

	t.Run("returns a list of verbs", func(t *testing.T) {
		t.Parallel()

		got := words.Verbs()

		if len(got) == 0 {
			t.Error("expected a list of verbs")
		}

		if !slices.Contains(got, "accomplish") {
			t.Error("expected the list to include the word 'accomplish'")
		}
	})
}

func TestNouns(t *testing.T) {
	t.Parallel()

	t.Run("returns a list of nouns", func(t *testing.T) {
		t.Parallel()

		got := words.Nouns()

		if len(got) == 0 {
			t.Error("expected a list of nouns")
		}

		if !slices.Contains(got, "absolute") {
			t.Error("expected the list to include the word 'absolute'")
		}
	})
}
