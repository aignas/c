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

func (f *frame) First() int {
	if len(f.vals) == 0 {
		return 0
	}
	return f.vals[0]
}

// Score returns the bowling score.
func Score(sheet string) (int, error) {
	var result int
	frames, err := parse(sheet)
	if err != nil {
		return 0, err
	}
	for _, frame := range frames {
		result += frame.First()
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
			frames[i].vals = append(frames[i].vals, 0)
		case '/':
			frames[i].vals = append(frames[i].vals, 10-frames[i].First())
		default:
			v, err := strconv.Atoi(string(throw))
			if err != nil {
				return nil, fmt.Errorf("bad input: %q", sheet)
			}
			frames[i].vals = append(frames[i].vals, v)
		}
		if i != 9 && len(frames[i].vals) == 2 {
			i++
		}
	}
	return frames, nil
}
