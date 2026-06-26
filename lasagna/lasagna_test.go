package lasagna

import "testing"

func TestOvenTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Calculates how many minutes the lasagna should be in the oven",
		layers:   0,
		time:     40,
		expected: 40,
	}

	if got := OvenTime; got != tc.expected {
		t.Errorf("OvenTime()  = %d; want %d", got, tc.expected)
	}

}

func TestRemainingOvenTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Remaining minutes in oven after 15 min",
		layers:   0,
		time:     15,
		expected: 25,
	}

	if got := RemainingOvenTime(tc.time); got != tc.expected {
		t.Errorf("RemainingOvenTime(%d) = %d; want %d", tc.time, got, tc.expected)
	}

}

func TestRemainingOvenTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Remaining minutes in oven after 30 min",
		layers:   0,
		time:     30,
		expected: 10,
	}

	if got := RemainingOvenTime(tc.time); got != tc.expected {
		t.Errorf("RemainingOvenTime(%d) = %d; want %d", tc.time, got, tc.expected)
	}

}

func TestPreparationTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Preparation time in minutes for one layer",
		layers:   1,
		time:     0,
		expected: 2,
	}

	if got := PreparationTime(tc.layers); got != tc.expected {
		t.Errorf("PreparationTime(%d) = %d; want %d", tc.layers, got, tc.expected)
	}

}

func TestPreparationTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Preparation time in minutes for multiple layers",
		layers:   4,
		time:     0,
		expected: 8,
	}

	if got := PreparationTime(tc.layers); got != tc.expected {
		t.Errorf("PreparationTime(%d) = %d; want %d", tc.layers, got, tc.expected)
	}

}

unc TestElapsedTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Total time in minutes for one layer",
		layers:   1,
		time:     30,
		expected: 32,
	}

	if got := ElapsedTime(tc.layers, tc.time); got != tc.expected {
		t.Errorf("ElapsedTime(%d, %d) = %d; want %d", tc.layers, tc.time, got, tc.expected)
	}

}

func TestElapsedTime(t *testing.T) {
	tc := lasagnaTests{
		name:     "Total time in minutes for multiple layers",
		layers:   4,
		time:     8,
		expected: 16,
	}

	if got := ElapsedTime(tc.layers, tc.time); got != tc.expected {
		t.Errorf("ElapsedTime(%d, %d) = %d; want %d", tc.layers, tc.time, got, tc.expected)
	}

}