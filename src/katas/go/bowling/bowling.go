package bowling

import (
	"fmt"
	"strconv"
)

func newFrame(vals ...int) frame {
	return frame{
		vals: vals,
	}
}

type frame struct {
	vals []int
}

func (f *frame) Sum() int {
	var r int
	for _, v := range f.vals {
		r += v
	}
	return r
}

func (f *frame) isSpare() bool {
	return len(f.vals) == 2 && f.Sum() == 10
}

func (f *frame) isStrike() bool {
	return len(f.vals) == 1 && f.First() == 10
}

func (f *frame) First() int {
	return f.vals[0]
}

// Result calculates the result of a bowling game.
func Result(input string) (int, error) {
	frames, err := parseFrames(input)
	if err != nil {
		return 0, fmt.Errorf("parsing frames: %s", err)
	}

	var result int
	for i, frame := range frames {
		result += frame.Sum()

		if frame.isSpare() {
			result += frames[i+1].First()
		}
		if frame.isStrike() {
			next := frames[i+1]
			result += next.First()
			if next.isStrike() {
				result += frames[i+2].First()
			} else {
				result += next.vals[1]
			}
		}
	}
	return result, nil
}

func parseFrames(input string) ([]frame, error) {
	var (
		frames = make([]frame, 10)
		i      = 0
		j      = 0
	)

	for _, r := range input {
		var val int

		switch r {
		case '-':
			// nothing
		case '/':
			val = 10 - frames[i].vals[j-1]
		case 'X':
			val = 10
		default:
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("bad input: %q", input)
			}
			val = v
		}

		frames[i].vals = append(frames[i].vals, val)
		if i == 9 || j == 0 && val != 10 {
			j++
			continue
		}

		i++
		j = 0
	}
	return frames, nil
}
