package username

import (
	"errors"
	"testing"
)

func TestGenerateWords(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name           string
		format         string
		expectedLength int
		expectedError  error
	}{
		{
			name:           "empty",
			format:         "",
			expectedLength: 0,
			expectedError:  ErrInvalidFormat,
		},
		{
			name:           "invalid format",
			format:         "{{invalid}}",
			expectedLength: 0,
			expectedError:  ErrInvalidFormat,
		},
		{
			name:           "valid format with all placeholders",
			format:         "{{adjective}} {{verb}} {{noun}} {{number}}",
			expectedLength: 4,
			expectedError:  nil,
		},
		{
			name:           "valid format with only two placeholders",
			format:         "{{adjective}}{{noun}}",
			expectedLength: 2,
			expectedError:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := generateWords(tt.format)
			if !errors.Is(err, tt.expectedError) {
				t.Fatalf("expected error %v, got %v", tt.expectedError, err)
			}

			if len(got) != tt.expectedLength {
				t.Fatalf("expected length %d, got %d", tt.expectedLength, len(got))
			}
		})
	}
}
