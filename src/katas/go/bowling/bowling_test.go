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
		{"-/-/-/-/-/-/-/-/-/-/-", 10 * 10, ""},
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
		{strings.Repeat("--", 10), repeatFrame(newFrame(0, 0), 10), ""},
		{strings.Repeat("3-", 10), repeatFrame(newFrame(3, 0), 10), ""},
		{strings.Repeat("45", 10), repeatFrame(newFrame(4, 5), 10), ""},
		{strings.Repeat("-/", 10) + "-", append(repeatFrame(newFrame(0, 10), 9), newFrame(0, 10, 0)), ""},
		{strings.Repeat("X", 12), append(repeatFrame(newFrame(10), 9), newFrame(10, 10, 10)), ""},
		{strings.Repeat("X", 10) + "-/", append(repeatFrame(newFrame(10), 9), newFrame(10, 0, 10)), ""},
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

func Test_newFrame(t *testing.T) {
	want := frame([]int{4, 2})
	assert.Equal(t, want, newFrame(4, 2))
	assert.Equal(t, 4, want.First())
	assert.Equal(t, 2, want.Last())
	assert.Len(t, want, 2)

	var empty frame
	assert.Equal(t, 0, empty.First())
	assert.Equal(t, 0, empty.Last())
	assert.Len(t, empty, 0)

	appended := empty.Append(4, 2)
	assert.Equal(t, want, appended)
	assert.Empty(t, empty, "should be unchanged")
}
