package bowling

import "strconv"

func newFrame(first int) frame {
	return frame{
		First: first,
	}
}

type frame struct {
	First int
}

// Score returns the bowling score.
func Score(sheet string) (int, error) {
	var result int
	frames, _ := parse(sheet)
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
		default:
			v, _ := strconv.Atoi(string(throw))
			frames[i].First = v
			i++
		}
	}
	return frames, nil
}
