package validate_test

import (
	"testing"

	"git.sr.ht/~jamesponddotco/randuser/internal/validate"
)

func TestFormat(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "valid input",
			input: "{{adjective}} {{noun}}",
			want:  true,
		},
		{
			name:  "invalid input with brackets",
			input: "{{random {{noun",
			want:  false,
		},
		{
			name:  "invalid input without brackets",
			input: "adjective noun",
			want:  false,
		},
		{
			name:  "invalid input with multiple brackets",
			input: "{{random}} {{random}}",
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := validate.Format(tt.input)
			if got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
