package bowling

type frame struct{}

func (*frame) Sum() int {
	return 0
}

// Score returns bowling score
func Score(input string) (int, error) {
	frames, err := parse(input)
	if err != nil {
		return 0, err
	}

	var result int
	for _, frame := range frames {
		result += frame.Sum()

		// we need to add extras
	}

	return result, nil
}

// parse returns frames
func parse(string) ([]frame, error) {
	return nil, nil
}
