package main

import "testing"

func Test_distance(t *testing.T) {
	tests := []struct {
		name         string
		fuelConsumed int
		fuelInTheTank int
		want         int
	}{
		{"You can ride for more than 150 kilometres", 7, 14, 200},
		{"You can ride for 100 kilometres", 9, 9, 100},
		{"You will ride for 10 kilometres", 10, 1, 10},
	}
	for _, test := range tests {
		got := distance(test.fuelConsumed, test.fuelInTheTank)
		if got != test.want {
			t.Error("Test fail:", test.name, "got:", got, "want:", test.want)
		}
	}
}