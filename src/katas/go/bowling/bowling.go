package bowling

import (
	"fmt"
	"strconv"
)

func newFrame(vals ...int) frame {
	return frame{
		First: vals[0],
	}
}

type frame struct {
	First int
}

// Score returns the bowling score.
func Score(sheet string) (int, error) {
	var result int
	frames, err := parse(sheet)
	if err != nil {
		return 0, err
	}
	for _, frame := range frames {
		result += frame.First
	}
	return result, nil
}

func parse(sheet string) ([]frame, error) {
	var (
		frames = make([]frame, 10)
		i      int
	)
	for _, throw := range sheet {
		switch throw {
		case '-':
			// nothing
		case '/':
			// nothing
		default:
			v, err := strconv.Atoi(string(throw))
			if err != nil {
				return nil, fmt.Errorf("bad input: %q", sheet)
			}
			frames[i].First = v
			i++
		}
	}
	return frames, nil
}
