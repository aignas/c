package bowling

import (
	"errors"
	"fmt"
	"strconv"
)

// _max is the number of pins
const (
	_max             = 10
	_frameCount      = 10
	_strikeFrameSize = 1
	_frameSize       = 2
)

// the length of the frame is either 2 or 3 and is enforced by the constructor
type frame []int

// newFrame returns the frame and the number of throws in this frame
func newFrame(s []int) (frame, error) {
	if len(s) != 2 && len(s) != 3 {
		return nil, fmt.Errorf("input must be 2 or 3 throws, got %d", len(s))
	}

	return frame(s), nil
}

func (f frame) Score() int {
	score := f[0] + f[1]
	if score >= _max {
		score += f[2]
	}

	return score
}

func (f frame) Size() int {
	if f[0] == _max {
		return _strikeFrameSize
	}

	return _frameSize
}

// Score returns bowling score
func Score(input string) (int, error) {
	throws, err := parse(input)
	if err != nil {
		return 0, err
	}

	frames, err := frames(throws)
	if err != nil {
		return 0, err
	}

	return sum(frames), nil
}

func sum(frames []frame) int {
	var result int

	for _, frame := range frames {
		result += frame.Score()
	}

	return result
}

// frames returns frames
func frames(throws []int) ([]frame, error) {
	frames := make([]frame, 0, _frameCount)

	for len(throws) != 0 {
		var input []int
		if len(throws) > _frameSize+1 {
			input = throws[:_frameSize+1]
		} else {
			input = throws
		}

		f, err := newFrame(input)
		if err != nil {
			return nil, err
		}

		frames = append(frames, f)

		switch {
		case len(frames) == _frameCount && f.Score() > _max:
			throws = throws[_frameSize+1:]
		case len(frames) == _frameCount:
			throws = throws[_frameSize:]
		default:
			throws = throws[f.Size():]
		}
	}

	if len(frames) < _frameCount {
		return nil, errors.New("too few throws")
	}

	if len(frames) > _frameCount {
		return nil, errors.New("too many throws")
	}

	return frames, nil
}

func parse(input string) ([]int, error) {
	result := make([]int, len(input))

	var first bool
	for i, r := range input {
		first = !first

		switch r {
		case '-':
			// nothing
		case '/':
			if first {
				return nil, errors.New("no spare on first throw")
			}

			result[i] = _max - result[i-1]
		case 'X':
			result[i] = _max
			first = false
		default:
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("invalid input: %q", r)
			}

			result[i] = v
		}
	}

	return result, nil
}
