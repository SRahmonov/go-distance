package main

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name            string
		volumeFuel      int
		consumptionFuel int
		want            int
	}{
		// TODO: Add test cases.
		{"More fuel in the tank than fuel", 9, 15, 166},
		{"fuel in the tank less than consumption", 9, 1, 11},
		{"fuel in the tank as much, as much fuel", 9, 9, 100},
	}
	for _, test := range tests {
		got := distance(test.volumeFuel, test.consumptionFuel)
		if got != test.want {
			t.Error("Test fail:", test.name, "got:", got, "want:", test.want)
		}

	}
}
