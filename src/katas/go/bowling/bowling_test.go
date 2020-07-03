package bowling

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input   string
		want    int
		wantErr string
	}{
		{"--------------------", 0, ""},
		{"5-5-5-5-5-5-5-5-5-5-", 50, ""},
		{"9-9-9-9-9-9-9-9-9-9-", 90, ""},
		{"5/5/5/5/5/5/5/5/5/5/5", 150, ""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := Result(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestParseFrames(t *testing.T) {
	wantAll := func(f frame) []frame {
		res := make([]frame, 10)
		for i := range res {
			res[i] = f
		}
		return res
	}
	wantAllWithLast := func(f, last frame) []frame {
		res := make([]frame, 10)
		for i := range res {
			if i == 9 {
				res[i] = last
			} else {
				res[i] = f
			}
		}
		return res
	}

	all := func(f string) string {
		res := make([]string, 10)
		for i := range res {
			res[i] = f
		}
		return strings.Join(res, "")
	}

	t.Parallel()
	tests := []struct {
		input   string
		want    []frame
		wantErr string
	}{
		{all("--"), wantAll(newFrame(0, 0)), ""},
		{all("5-"), wantAll(newFrame(5, 0)), ""},
		{all("54"), wantAll(newFrame(5, 4)), ""},
		{all("5/") + "5", wantAllWithLast(newFrame(5, 5), newLastFrame(5, 5, 5)), ""},
		{"foo", nil, "bad input: \"foo\""},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := parseFrames(tt.input)
			assert.Equal(t, tt.want, got)
			if tt.wantErr != "" {
				assert.EqualError(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_frame_sum(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg   string
		frame frame
		want  int
	}{
		{msg: "empty"},
		{msg: "5-", frame: newFrame(5, 0), want: 5},
		{msg: "5-", frame: newFrame(0, 4), want: 4},
		{msg: "-/4", frame: newLastFrame(0, 10, 4), want: 14},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.frame.Sum())
		})
	}
}

func Test_frame_isSpare(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg   string
		frame frame
		want  bool
	}{
		{msg: "5-", frame: newFrame(5, 0), want: false},
		{msg: "5/", frame: newFrame(5, 5), want: true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.frame.isSpare())
		})
	}
}
