package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n cannot be less than 1")
	}
	cntSteps := 0
	for ; n != 1; cntSteps++ {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return cntSteps, nil
}
