package bowling

import (
	"fmt"
	"strconv"
)

func newFrame(first, second int) frame {
	return frame{
		First:  first,
		Second: second,
	}
}

func newLastFrame(first, second, third int) frame {
	return frame{
		First:  first,
		Second: second,
		Third:  third,
	}
}

type frame struct {
	First  int
	Second int
	Third  int
}

func (f *frame) Sum() int {
	return f.First + f.Second + f.Third
}

func (f *frame) isSpare() bool {
	return f.Sum() == 10
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

		if i != 9 && frame.isSpare() {
			result += frames[i+1].First
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
			val = 10 - frames[i].First
		default:
			v, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("bad input: %q", input)
			}
			val = v
		}

		if j == 0 {
			frames[i].First = val
			j++
		} else if j == 1 {
			frames[i].Second = val
			if i == 9 {
				j++
			} else {
				j = 0
				i++
			}
		} else {
			frames[i].Third = val
		}
	}
	return frames, nil
}
