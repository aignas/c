package bowling

import "strconv"

type frame struct{}

// Score returns the bowling score.
func Score(sheet string) (int, error) {
	var result int
	for _, throw := range sheet {
		switch throw {
		case '-':
			// nothing
		default:
			v, _ := strconv.Atoi(string(throw))
			result += v
		}
	}
	return result, nil
}

func parse(string) ([]frame, error) {
	return make([]frame, 10), nil
}
