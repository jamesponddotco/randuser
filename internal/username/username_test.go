package username_test

import (
	"errors"
	"testing"

	"git.sr.ht/~jamesponddotco/randuser/internal/username"
)

func TestGenerate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		format        string
		expectedError error
	}{
		{
			name:          "empty",
			format:        "",
			expectedError: username.ErrInvalidFormat,
		},
		{
			name:          "invalid format",
			format:        "{{invalid}}",
			expectedError: username.ErrInvalidFormat,
		},
		{
			name:          "valid format with all placeholders",
			format:        "{{adjective}} {{verb}} {{noun}} {{number}}",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := username.Generate(tt.format)
			if !errors.Is(err, tt.expectedError) {
				t.Fatalf("expected error %v, got %v", tt.expectedError, err)
			}

			if tt.expectedError == nil {
				if got == "" {
					t.Fatalf("expected non-empty username, got empty")
				}
			}
		})
	}
}

func TestGenerateWithTitleCase(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		format        string
		expectedError error
	}{
		{
			name:          "empty",
			format:        "",
			expectedError: username.ErrInvalidFormat,
		},
		{
			name:          "invalid format",
			format:        "{{invalid}}",
			expectedError: username.ErrInvalidFormat,
		},
		{
			name:          "valid format with all placeholders",
			format:        "{{adjective}} {{verb}} {{noun}} {{number}}",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := username.GenerateWithTitleCase(tt.format)
			if !errors.Is(err, tt.expectedError) {
				t.Fatalf("expected error %v, got %v", tt.expectedError, err)
			}

			if tt.expectedError == nil {
				if got == "" {
					t.Fatalf("expected non-empty username, got empty")
				}
			}
		})
	}
}
