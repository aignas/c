package bowling

import (
	"fmt"
	"strconv"
)

// _max is the number of pins that we can hit at most during one throw.
const _max = 10

func newFrame(vals ...int) frame {
	return frame(vals)
}

type frame []int

func (f frame) First() int {
	if len(f) == 0 {
		return 0
	}

	return f[0]
}

func (f frame) Last() int {
	if len(f) == 0 {
		return 0
	}

	return f[len(f)-1]
}

func (f frame) Append(vals ...int) frame {
	return append(f, vals...)
}

// Score returns the bowling score.
func Score(sheet string) (int, error) {
	var result int

	frames, err := parse(sheet)
	if err != nil {
		return 0, err
	}

	for _, frame := range frames {
		result += frame.First() + frame[1]
	}

	return result, nil
}

func parse(sheet string) ([]frame, error) {
	var (
		frames = make([]frame, 10)
		i      int
	)

	for _, throw := range sheet {
		var val int

		switch throw {
		case '-':
			val = 0
		case 'X':
			val = _max
		case '/':
			val = _max - frames[i].Last()
		default:
			v, err := strconv.Atoi(string(throw))
			if err != nil {
				return nil, fmt.Errorf("bad input: %q", sheet)
			}

			val = v
		}

		frames[i] = frames[i].Append(val)
		if i != 9 && (len(frames[i]) == 2 || val == _max) {
			i++
		}
	}
	return frames, nil
}
