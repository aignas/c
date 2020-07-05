package bowling

type frame []int

func (frame) Score() int {
	return 0
}

func (frame) Len() int {
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
		result += frame.Score()
	}

	return result, nil
}

// parse returns frames
func parse(input string) ([]frame, error) {
	throws, err := parseThrows(input)
	if err != nil {
		return nil, err
	}

	var j int
	frames := make([]frame, 10)
	for i := range frames {
		if i == 9 {
			frames[i] = newFrame(throws[j:])
		} else {
			frames[i] = newFrame(throws[j : j+3])
		}

		j += frames[i].Len()
	}

	return nil, nil
}

func parseThrows(input string) ([]int, error) {
	return make([]int, len(input)), nil
}

// newFrame returns the frame and the number of throws in this frame
func newFrame(s []int) frame {
	return nil
}
