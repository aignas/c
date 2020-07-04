package bowling

import "strconv"

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
