package bowling

import (
	"errors"
	"fmt"
	"strconv"
)

// _max is the number of pins
const (
	_max             = 10
	_strikeFrameSize = 1
	_frameSize       = 2
)

// the length of the frame is either 2 or 3 and is enforced by the constructor
type frame []int

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

	return sum(frames(throws)), nil
}

func sum(frames []frame) int {
	var result int

	for _, frame := range frames {
		result += frame.Score()
	}

	return result
}

// frames returns frames
func frames(throws []int) []frame {
	var j int

	frames := make([]frame, 10)
	for i := range frames {
		if i == len(frames)-1 {
			frames[i], _ = newFrame(throws[j:])
		} else {
			frames[i], _ = newFrame(throws[j : j+3])
		}

		j += frames[i].Size()
	}

	return frames
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

// newFrame returns the frame and the number of throws in this frame
func newFrame(s []int) (frame, error) {
	if len(s) != 2 && len(s) != 3 {
		return nil, errors.New("input must be 2 or 3 throws")
	}

	return frame(s), nil
}
