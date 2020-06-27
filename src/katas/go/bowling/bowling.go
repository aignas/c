package bowling

import (
	"fmt"
	"strconv"
)

type (
	frame struct {
		First  int
		Second int
		Extra  int
	}
)

func newFrame(vals ...int) frame {
	if len(vals) == 2 {
		return frame{
			First:  vals[0],
			Second: vals[1],
		}
	}
	return frame{
		First:  vals[0],
		Second: vals[1],
		Extra:  vals[2],
	}
}

// Result calculates the bowling score
func Result(in string) int {
	var result int
	frames, _ := parseFrames(in)
	for i, f := range frames {
		result += f.First + f.Second + f.Extra

		if i == 9 {
			continue
		}

		next := frames[i+1]

		if f.IsSpare() {
			result += next.First
			continue
		}

		if !f.IsStrike() {
			continue
		}

		if next.IsStrike() {
			result += 10 + frames[i+2].First
			continue
		}
		result += next.First + next.Second
	}
	return result
}

func parseFrames(in string) ([]frame, error) {
	var (
		result []frame
		vals   []int
	)
	for _, r := range in {
		if len(vals) == 2 && len(result) != 9 {
			result = append(result, newFrame(vals...))
			vals = nil
		}

		switch r {
		case '-':
			vals = append(vals, 0)
		case '/':
			if len(vals) == 0 {
				return nil, fmt.Errorf("spare is impossible on the first throw")
			}
			vals = append(vals, 10-vals[0])
		case 'X':
			if len(result) == 9 {
				vals = append(vals, 10)
			} else {
				vals = append(vals, 10, 0)
			}
		default:
			val, err := strconv.Atoi(string(r))
			if err != nil {
				return nil, fmt.Errorf("failed to parse: %s", err)
			}
			vals = append(vals, val)
		}
	}
	return append(result, newFrame(vals...)), nil
}

func (f *frame) IsSpare() bool {
	return !f.IsStrike() && f.First+f.Second == 10
}

func (f *frame) IsStrike() bool {
	return f.First == 10 && f.Second == 0
}
