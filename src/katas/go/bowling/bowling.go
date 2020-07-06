package bowling

import (
	"fmt"
	"strconv"
)

type frame []int

func (frame) Score() int {
	return 0
}

func (frame) Size() int {
	return 0
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
			frames[i] = newFrame(throws[j:])
		} else {
			frames[i] = newFrame(throws[j : j+3])
		}

		j += frames[i].Size()
	}

	return frames
}

func parse(input string) ([]int, error) {
	result := make([]int, len(input))
	for i, r := range input {
		var val int
		switch r {
		case '-':
			// nothing
		case '/':
			val = 10 - result[i-1]
		case 'X':
			val = 10
		default:
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("invalid input: %q", r)
			}
			val = v
		}
		result[i] = val
	}
	return result, nil
}

// newFrame returns the frame and the number of throws in this frame
func newFrame(s []int) frame {
	return nil
}
