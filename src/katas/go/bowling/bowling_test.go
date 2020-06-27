package bowling

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  int
	}{
		{"9-9-9-9-9-9-9-9-9-9-", 9 * 10},
		{"5-5-5-5-5-5-5-5-5-5-", 5 * 10},
		{"XXXXXXXXXXXX", 300},
		{"------------------XXX", 30},
		{"5/5/5/5/5/5/5/5/5/5/5", 150},
		{"X5/1---------------", 20 + 11 + 1},
		{"X5-1---------------", 15 + 5 + 1},
		{"XX1---------------", 21 + 11 + 1},
		{"XXX1--------------", 30 + 21 + 11 + 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.want, Result(tt.input))
		})
	}
}

type frameTest struct {
	msg     string
	input   string
	want    []frame
	wantErr string
}

func Test_parseFrames(t *testing.T) {
	bad := func(in, err string) frameTest {
		return frameTest{
			msg:     "err: " + in,
			input:   in,
			wantErr: err,
		}
	}
	ok := func(in string, want ...frame) frameTest {
		return frameTest{
			msg:   "ok: " + in,
			input: in,
			want:  want,
		}
	}

	strike := frame{First: 10}
	none := frame{}

	t.Parallel()
	tests := []frameTest{
		bad("/", "spare is impossible on the first throw"),
		ok("9-", frame{First: 9}),
		ok("5-", frame{First: 5}),
		ok("-5", frame{Second: 5}),
		ok("-/", frame{Second: 10}),
		ok("5/", frame{First: 5, Second: 5}),
		ok("X", strike),
		ok("XX", strike, strike),
		ok("XXXXXXXXXXXX", strike, strike, strike, strike, strike, strike, strike, strike, strike, frame{First: 10, Second: 10, Extra: 10}),
		ok("------------------XXX", none, none, none, none, none, none, none, none, none, frame{First: 10, Second: 10, Extra: 10}),
		ok("------------------X--", none, none, none, none, none, none, none, none, none, frame{First: 10}),
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
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

func Test_frame_IsStrike(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg   string
		frame frame
		want  bool
	}{
		{"strike", frame{First: 10}, true},
		{"not", frame{First: 5}, false},
		{"last not", frame{First: 10, Second: 1}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.frame.IsStrike())
		})
	}
}

func Test_frame_IsSpare(t *testing.T) {
	t.Parallel()
	tests := []struct {
		msg   string
		frame frame
		want  bool
	}{
		{"strike", frame{First: 10}, false},
		{"spare", frame{Second: 10}, true},
		{"spare2", frame{First: 5, Second: 5}, true},
		{"not", frame{First: 5, Second: 4}, false},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.msg, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, tt.frame.IsSpare())
		})
	}
}
