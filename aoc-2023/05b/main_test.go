package main

import "testing"

func Test_gardenmap_Convert(t *testing.T) {

	seed_to_soil := gardenmap{
		Name: "seed-to-soil",
		Rules: &[]convertrule{
			{srcStart: 98, destStart: 50, rangeLength: 2},
			{srcStart: 50, destStart: 52, rangeLength: 48},
		}}

	tests := []struct {
		name string
		gardenmap
		input int
		want  int
	}{
		{
			name:      "seed number 10 corresponds to soil number 10",
			gardenmap: seed_to_soil,
			input:     10,
			want:      10,
		},
		{
			name:      "seed number 53 corresponds to soil number 55",
			gardenmap: seed_to_soil,
			input:     53,
			want:      55,
		},
		{
			name:      "seed number 98 corresponds to soil number 50",
			gardenmap: seed_to_soil,
			input:     98,
			want:      50,
		},
		{
			name:      "and that seed number 99 corresponds to soil number 51",
			gardenmap: seed_to_soil,
			input:     99,
			want:      51,
		},
		{
			name:      "and that seed number 100 corresponds to soil number 100",
			gardenmap: seed_to_soil,
			input:     100,
			want:      100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gardenmap.Convert(tt.input); got != tt.want {
				t.Errorf("gardenmap.Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
