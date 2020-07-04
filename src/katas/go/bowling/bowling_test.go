package bowling

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input   string
		want    int
		wantErr string
	}{
		{"--------------------", 0, ""},
		{"5-5-5-5-5-5-5-5-5-5-", 5 * 10, ""},
		{"3-3-3-3-3-3-3-3-3-3-", 3 * 10, ""},
		{"bad-input", 0, "bad input: \"bad-input\""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := Score(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestParse(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input   string
		want    []frame
		wantErr string
	}{
		{strings.Repeat("--", 10), make([]frame, 10), ""},
		{strings.Repeat("3-", 10), repeatFrame(newFrame(3), 10), ""},
		{"bad-input", nil, "bad input: \"bad-input\""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := parse(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func repeatFrame(f frame, n int) []frame {
	frames := make([]frame, n)
	for i := range frames {
		frames[i] = f
	}
	return frames
}
