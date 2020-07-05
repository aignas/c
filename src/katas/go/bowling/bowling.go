package bowling

type frame struct{}

func (*frame) Sum() int {
	return 0
}

func (*frame) First() int {
	return 0
}

func (*frame) Head() int {
	return 0
}

func (*frame) IsSpare() bool {
	return false
}

func (*frame) IsStrike() bool {
	return false
}

// Score returns bowling score
func Score(input string) (int, error) {
	frames, err := parse(input)
	if err != nil {
		return 0, err
	}

	var (
		result        int
		doubleFirst   bool
		doubleNextTwo bool
	)

	for _, frame := range frames {
		result += frame.Sum()
		if doubleFirst {
			result += frame.First()
		}

		if doubleNextTwo {
			if frame.IsStrike() {
				result += frame.First()
			} else {
				result += frame.Head()
			}
		}

		// we need to add extras
		doubleFirst = frame.IsSpare() || doubleNextTwo && frame.IsStrike()
		doubleNextTwo = frame.IsStrike()
	}

	return result, nil
}

// parse returns frames
func parse(string) ([]frame, error) {
	return nil, nil
}
