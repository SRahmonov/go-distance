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
		{"Mercedes-Benz", 8, 5, 62},
		{"Audi-Q7", 9, 7, 77},
		{"BMW-X7", 9, 8, 88},
	}
	for _, test := range tests {
		got := distance(test.volumeFuel, test.consumptionFuel)
		if got != test.want {
			t.Error("Test fail:", test.name, "got:", got, "want:", test.want)
		}

	}
}
