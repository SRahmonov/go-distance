package main

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name string
		volumeFuel  float32
		expenseFuel float32
		want float32
	}{
		// TODO: Add test cases.
		{"Mercedes-Benz", 8.5, 7.5,88.2352941 },
		{"Audi-Q7", 9.8, 8.2,83.6734694},
		{"BMW-X7",9.2, 8.6,93.4782609},
	}
	for _, test := range tests {
		got := distance(test.volumeFuel, test.expenseFuel)
		if got !=test.want{
			t.Error("Test fail:", test.name, "got:", test.want)
		}

	}
}