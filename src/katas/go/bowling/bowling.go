package bowling

type frame []int

func (frame) Score() int {
	return 0
}

func (frame) Size() int {
	return 0
}

// Score returns bowling score
func Score(input string) (int, error) {
	throws, err := parseThrows(input)
	if err != nil {
		return 0, err
	}

	return sum(frames(throws)), nil
}

func sum(frames []frame) int {
	var result int

	for _, frame := range frames {
		result += frame.Score()
	}

	return result
}

// frames returns frames
func frames(throws []int) []frame {
	var j int

	frames := make([]frame, 10)
	for i := range frames {
		if i == len(frames)-1 {
			frames[i] = newFrame(throws[j:])
		} else {
			frames[i] = newFrame(throws[j : j+3])
		}

		j += frames[i].Size()
	}

	return frames
}

func parseThrows(input string) ([]int, error) {
	return make([]int, len(input)), nil
}

// newFrame returns the frame and the number of throws in this frame
func newFrame(s []int) frame {
	return nil
}
