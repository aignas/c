package bowling

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScore(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input   string
		want    int
		wantErr string
	}{
		{strings.Repeat("-", 20), 0, ""},
		{strings.Repeat("5-", 10), 50, ""},
		{strings.Repeat("5/", 10) + "5", 150, ""},
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
		want    []int
		wantErr string
	}{
		{"--", []int{0, 0}, ""},
		{"45", []int{4, 5}, ""},
		{"4/", []int{4, 6}, ""},
		{"X", []int{10}, ""},
		{"/", nil, "no spare on first throw"},
		{"X/", nil, "no spare on first throw"},
		{"45/", nil, "no spare on first throw"},
		{"4y", nil, "invalid input: 'y'"},
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

func Test_newFrame(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg     string
		input   []int
		want    frame
		wantErr string
	}{
		{
			msg:   "ok if 2 empty throws",
			input: []int{0, 0},
			want:  frame([]int{0, 0}),
		},
		{
			msg:   "ok if 3 empty throws",
			input: []int{0, 0, 0},
			want:  frame([]int{0, 0, 0}),
		},
		{
			msg:     "err if size is less than 2",
			input:   []int{10},
			wantErr: "input must be 2 or 3 throws",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()

			got, err := newFrame(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_frame_Size(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg    string
		throws []int
		want   int
	}{
		{"simple", []int{4, 5, 4}, 2},
		{"strike", []int{10, 5, 4}, 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()

			frame, err := newFrame(tt.throws)
			require.NoError(t, err)

			got := frame.Size()

			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_frame_Score(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg    string
		throws []int
		want   int
	}{
		{"simple", []int{4, 5, 4}, 9},
		{"spare", []int{4, 6, 4}, 14},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()

			frame, err := newFrame(tt.throws)
			require.NoError(t, err)

			got := frame.Score()

			assert.Equal(t, tt.want, got)
		})
	}
}
